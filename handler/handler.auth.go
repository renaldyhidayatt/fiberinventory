package handler

import (
	"fiberinventory/entity"
	"fiberinventory/pkg"
	"fiberinventory/schemas"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
)

type handlerAuth struct {
	auth entity.EntityUser
}

func NewHandlerAuth(auth entity.EntityUser) *handlerAuth {
	return &handlerAuth{auth: auth}
}

func (h *handlerAuth) HandlerHello(c *fiber.Ctx) error {
	return c.SendString("Handler Auth")
}

func (h *handlerAuth) HandlerRegister(c *fiber.Ctx) error {
	var body schemas.SchemaUser

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.auth.EntityRegister(&body)

	if error.Type == "error_register_01" {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"msg": "Email already taken",
		})

	}
	if error.Type == "error_register_02" {
		return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
			"msg": "Register new user account failed",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"error": false,
		"msg":   "Register new user account success",
	})
}

func (h *handlerAuth) HandlerLogin(c *fiber.Ctx) error {
	var body schemas.SchemaUser

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	res, error := h.auth.EntityLogin(&body)

	if error.Type == "error_login_01" {
		return c.Status(fiber.StatusExpectationFailed).JSON(fiber.Map{
			"msg": "User account is not never registered",
		})
	}

	if error.Type == "error_login_02" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg": "Email or password is wrong",
		})
	}

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["email"] = res.Email
	claims["user_id"] = res.ID
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString([]byte(pkg.GodotEnv("JWT_SECRET")))
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"status":  "success",
		"message": "Login susccses",
		"data":    t,
	})

}

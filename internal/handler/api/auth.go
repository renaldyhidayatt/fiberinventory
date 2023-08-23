package api

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/pb"

	"github.com/gofiber/fiber/v2"
)

type handleAuth struct {
	client pb.AuthServiceClient
}

func NewHandlerAuth(client pb.AuthServiceClient, router *fiber.App) {
	handlerAuth := handleAuth{
		client: client,
	}

	routerAuth := router.Group("/api/auth")

	routerAuth.Get("/hello", handlerAuth.handlerHello)
	routerAuth.Post("/register", handlerAuth.register)
	routerAuth.Post("/login", handlerAuth.login)
}

// handlerHello function
// @Summary Greet the user
// @Description Return a greeting message to the user
// @Tags Auth
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /auth/hello [get]
func (h *handleAuth) handlerHello(c *fiber.Ctx) error {
	return c.SendString("Handler Auth")
}

// register function
// @Summary Register to the application
// @Description Create User
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body domain.RegisterInput true "User information"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /auth/register [post]
func (h *handleAuth) register(c *fiber.Ctx) error {
	var body domain.RegisterInput

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	data := &pb.SignUpUserInput{
		Firstname: body.FirstName,
		Lastname:  body.LastName,
		Email:     body.Email,
		Password:  body.Password,
		Role:      body.Role,
	}

	res, err := h.client.RegisterUser(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil register",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// login function
// @Summary Login to the application
// @Description Login with username and password to get a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body domain.LoginInput true "User information"
// @Success 200 {object} domain.ResponseAuth
// @Failure 400 {object} domain.ErrorMessage
// @Router /auth/login [post]
func (h *handleAuth) login(c *fiber.Ctx) error {
	var body domain.LoginInput

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	data := &pb.SignInUserInput{
		Email:    body.Email,
		Password: body.Password,
	}

	res, err := h.client.LoginUser(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.ResponseAuth{
		Message:    "Success login",
		StatusCode: fiber.StatusOK,
		Jwt:        res,
	})

}

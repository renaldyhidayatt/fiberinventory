package v1

import (
	"fiberinventory/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initAuthGroup(api *fiber.App) {
	users := api.Group("/auth")

	users.Get("/hello", h.handlerHello)
	users.Post("/register", h.register)
	users.Post("/login", h.login)
}

// handlerHello function
// @Summary Greet the user
// @Description Return a greeting message to the user
// @Tags Auth
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /auth/hello [get]
func (h *Handler) handlerHello(c *fiber.Ctx) error {
	return c.SendString("Handler Auth")
}

// register function
// @Summary Register to the application
// @Description Create User
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body domain.UserInput true "User information"
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /auth/register [post]
func (h *Handler) register(c *fiber.Ctx) error {
	var body domain.UserInput

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

	_, err := h.services.User.Register(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil register",
		Data:       nil,
		StatusCode: fiber.StatusOK,
	})
}

// login function
// @Summary Login to the application
// @Description Login with username and password to get a JWT token
// @Tags Auth
// @Accept json
// @Produce json
// @Param user body domain.UserInput true "User information"
// @Success 200 {object} domain.ResponseAuth
// @Failure 400 {object} domain.ErrorMessage
// @Router /auth/login [post]
func (h *Handler) login(c *fiber.Ctx) error {
	var body domain.UserInput

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

	res, err := h.services.User.Login(&body)

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

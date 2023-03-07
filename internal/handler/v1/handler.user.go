package v1

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initUserGroup(api *fiber.App) {
	routerUser := api.Group("/users")

	routerUser.Use(middleware.Proctected())

	routerUser.Get("/hello", h.handlerUserHello)
	routerUser.Put("/update/:id", h.handlerUserUpdate)
	routerUser.Delete("/delete/:id", h.handlerUserDelete)
	routerUser.Get("/", h.handlerUserResults)
	routerUser.Get("/:id", h.handlerUserResult)
}

// handlerUserHello function
// @Summary Greet the Users
// @Description Return a greeting message to the users
// @Tags Users
// @Produce plain
// @Security BearerAuth
// @Success 200 {string} string "OK"
// @Router /users/hello [get]
func (h *Handler) handlerUserHello(c *fiber.Ctx) error {
	return c.SendString("Handler user")
}

// handlerUserResults function
// @Summary Get users results
// @Description Retrieve the results for each USer
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /users [get]
func (h *Handler) handlerUserResults(c *fiber.Ctx) error {
	res, err := h.services.User.Results()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil mendapatkan data",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// handlerUserResult function
// @Summary Get user result
// @Description Retrieve the result for a specific user
// @Tags Users
// @Param id path string true "User ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /users/{id} [get]
func (h *Handler) handlerUserResult(c *fiber.Ctx) error {
	var body domain.UserInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.User.Result(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil mendapatkan data",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// handlerUserDelete function
// @Summary Delete user
// @Description Delete a specific user
// @Tags Users
// @Param id path string true "User ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /users/{id} [delete]
func (h *Handler) handlerUserDelete(c *fiber.Ctx) error {
	var body domain.UserInput

	id := c.Params("id")

	body.ID = id

	res, err := h.services.User.Delete(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil menghapus user",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// handlerUserUpdate function
// @Summary Update User
// @Description Update a specific user
// @Tags Users
// @Param id path string true "Users ID"
// @Param body body domain.UserInput true "User Data"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /users/{id} [put]
func (h *Handler) handlerUserUpdate(c *fiber.Ctx) error {
	var body domain.UserInput

	id := c.Params("id")

	body.ID = id

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

	res, err := h.services.User.Update(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil mengupdate users",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

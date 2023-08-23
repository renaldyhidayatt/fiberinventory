package api

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/pb"

	"github.com/gofiber/fiber/v2"
)

type handlerUser struct {
	client pb.UserServiceClient
}

func NewHandlerUser(client pb.UserServiceClient, router *fiber.App) {
	h := handlerUser{
		client: client,
	}

	routerUser := router.Group("/api/user")

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
func (h *handlerUser) handlerUserHello(c *fiber.Ctx) error {
	return c.SendString("Handler user")
}

// handlerUserResults function
// @Summary Get users results
// @Description Retrieve the results for each User
// @Tags Users
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /users [get]
func (h *handlerUser) handlerUserResults(c *fiber.Ctx) error {
	res, err := h.client.GetUsers(c.Context(), &pb.UsersRequest{})

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
func (h *handlerUser) handlerUserResult(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.UserRequest{
		Id: id,
	}

	res, err := h.client.GetUser(c.Context(), data)

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
func (h *handlerUser) handlerUserDelete(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.UserRequest{
		Id: id,
	}

	res, err := h.client.DeleteUser(c.Context(), data)

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
// @Param id path string true "User ID"
// @Param body body domain.UpdateUserRequest true "User Data"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /users/{id} [put]
func (h *handlerUser) handlerUserUpdate(c *fiber.Ctx) error {
	var body domain.UpdateUserRequest

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

	data := &pb.UpdateUserInput{
		Id:        id,
		Firstname: body.FirstName,
		Lastname:  body.LastName,
		Email:     body.Email,
		Role:      body.Role,
	}

	res, err := h.client.UpdateUser(c.Context(), data)

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

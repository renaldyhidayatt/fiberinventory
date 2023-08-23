package api

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"
	"fiberinventory/internal/pb"

	"github.com/gofiber/fiber/v2"
)

type handleCustomer struct {
	client pb.CustomerServiceClient
}

func NewHandlerCustomer(client pb.CustomerServiceClient, router *fiber.App) {
	handlerCustomer := handleCustomer{
		client: client,
	}

	routerCustomer := router.Group("/api/customer")

	routerCustomer.Use(middleware.Proctected())
	routerCustomer.Get("/hello", handlerCustomer.handlerCustomerHello)
	routerCustomer.Get("/", handlerCustomer.handlerCustomerResults)
	routerCustomer.Get("/:id", handlerCustomer.handlerCustomerResult)
	routerCustomer.Post("/create", handlerCustomer.handlerCustomerCreate)
	routerCustomer.Post("/update/:id", handlerCustomer.handlerCustomerUpdate)
	routerCustomer.Post("/delete/:id", handlerCustomer.handlerCustomerDelete)

}

// handlerCustomerHello function
// @Summary Greet the Customer
// @Description Return a greeting message to the customer
// @Tags Customer
// @Produce plain
// @Security BearerAuth
// @Success 200 {string} string "OK"
// @Router /customer/hello [get]
func (h *handleCustomer) handlerCustomerHello(c *fiber.Ctx) error {
	return c.SendString("Handler customer")

}

// handlerCustomerCreate function
// @Summary createCustomer to the application
// @Description Create Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param body body domain.CreateCustomerRequest true "Customer information"
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /customer/create [post]
func (h *handleCustomer) handlerCustomerCreate(c *fiber.Ctx) error {
	var body domain.CreateCustomerRequest

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

	data := &pb.CreateCustomerInput{
		Name:    body.Name,
		Alamat:  body.Alamat,
		Email:   body.Email,
		Telepon: body.Telepon,
	}

	res, err := h.client.CreateCustomer(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil membuat customer",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// handlerCustomerResults function
// @Summary Get customer results
// @Description Retrieve the results for each customer
// @Tags Customer
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /api/customer [get]
func (h *handleCustomer) handlerCustomerResults(c *fiber.Ctx) error {
	res, err := h.client.GetCustomers(c.Context(), &pb.CustomersRequest{})

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

// handlerCustomerResult function
// @Summary Get customer result
// @Description Retrieve the result for a specific customer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /customer/{id} [get]
func (h *handleCustomer) handlerCustomerResult(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.CustomerRequest{
		Id: id,
	}

	res, err := h.client.GetCustomer(c.Context(), data)

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

// handlerCustomerUpdate function
// @Summary Update Customer
// @Description Update a specific customer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Param body body domain.UpdateCustomerRequest true "Customer Data"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /customer/{id} [put]
func (h *handleCustomer) handlerCustomerUpdate(c *fiber.Ctx) error {
	var body domain.UpdateCustomerRequest

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

	data := &pb.UpdateCustomerInput{
		Id:      body.ID,
		Name:    body.Name,
		Alamat:  body.Alamat,
		Email:   body.Email,
		Telepon: body.Telepon,
	}

	res, err := h.client.UpdateCustomer(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil mengupdate category",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// handlerCustomerDelete function
// @Summary Delete customer
// @Description Delete a specific customer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /customer/{id} [delete]
func (h *handleCustomer) handlerCustomerDelete(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.CustomerRequest{
		Id: id,
	}

	res, err := h.client.DeleteCustomer(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil menghapus category",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

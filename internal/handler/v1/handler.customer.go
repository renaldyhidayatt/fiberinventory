package v1

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initCustomer(api *fiber.App) {
	routerCustomer := api.Group("/customer")

	routerCustomer.Use(middleware.Proctected())

	routerCustomer.Get("/hello", h.handlerCustomerHello)
	routerCustomer.Get("/", h.handlerCustomerResults)
	routerCustomer.Get("/:id", h.handlerCustomerResult)
	routerCustomer.Post("/create", h.handlerCustomerCreate)
	routerCustomer.Post("/update/:id", h.handlerCustomerUpdate)
	routerCustomer.Post("/delete/:id", h.handlerCustomerDelete)

}

// handlerCustomerHello function
// @Summary Greet the Customer
// @Description Return a greeting message to the customer
// @Tags Customer
// @Produce plain
// @Security BearerAuth
// @Success 200 {string} string "OK"
// @Router /customer/hello [get]
func (h *Handler) handlerCustomerHello(c *fiber.Ctx) error {
	return c.SendString("Handler customer")

}

// handlerCustomerCreate function
// @Summary createCustomer to the application
// @Description Create Customer
// @Tags Customer
// @Accept json
// @Produce json
// @Param body body domain.CustomerInput true "Customer information"
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /customer/create [post]
func (h *Handler) handlerCustomerCreate(c *fiber.Ctx) error {
	var body domain.CustomerInput

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

	res, err := h.services.Customer.Create(&body)

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
// @Router /customer [get]
func (h *Handler) handlerCustomerResults(c *fiber.Ctx) error {
	res, err := h.services.Customer.Results()

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
func (h *Handler) handlerCustomerResult(c *fiber.Ctx) error {
	var body domain.CustomerInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Customer.Result(&body)

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
func (h *Handler) handlerCustomerDelete(c *fiber.Ctx) error {
	var body domain.CustomerInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Customer.Result(&body)

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

// handlerCustomerUpdate function
// @Summary Update Customer
// @Description Update a specific customer
// @Tags Customer
// @Param id path string true "Customer ID"
// @Param body body domain.CustomerInput true "Customer Data"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /customer/{id} [put]
func (h *Handler) handlerCustomerUpdate(c *fiber.Ctx) error {
	var body domain.CustomerInput
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

	res, err := h.services.Customer.Update(&body)

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

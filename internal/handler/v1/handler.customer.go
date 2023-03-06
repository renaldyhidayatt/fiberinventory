package v1

import (
	"fiberinventory/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initCustomer(api *fiber.App) {
	routerCustomer := api.Group("/customer")

	routerCustomer.Get("/hello", h.handlerCustomerHello)
	routerCustomer.Get("/", h.handlerCustomerResults)
	routerCustomer.Get("/:id", h.handlerCustomerResult)
	routerCustomer.Post("/create", h.handlerCustomerCreate)
	routerCustomer.Post("/update/:id", h.handlerCustomerUpdate)
	routerCustomer.Post("/delete/:id", h.handlerCustomerDelete)

}

func (h *Handler) handlerCustomerHello(c *fiber.Ctx) error {
	return c.SendString("Handler customer")

}

func (h *Handler) handlerCustomerCreate(c *fiber.Ctx) error {
	var body domain.CustomerInput

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	res, err := h.services.Customer.Create(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"data":  res,
		"msg":   "create successfully",
	})
}

func (h *Handler) handlerCustomerResults(c *fiber.Ctx) error {
	res, err := h.services.Customer.Results()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Customer data already to use",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerCustomerResult(c *fiber.Ctx) error {
	var body domain.CustomerInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Customer.Result(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Customer found",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerCustomerDelete(c *fiber.Ctx) error {
	var body domain.CustomerInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Customer.Result(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Category  found",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerCustomerUpdate(c *fiber.Ctx) error {
	var body domain.CustomerInput
	id := c.Params("id")

	body.ID = id

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, err := h.services.Customer.Update(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Update customer data success for this id",
		"code": fiber.StatusCreated,
	})
}

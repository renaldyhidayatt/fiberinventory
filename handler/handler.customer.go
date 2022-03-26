package handler

import (
	"fiberinventory/entity"
	"fiberinventory/schemas"

	"github.com/gofiber/fiber/v2"
)

type handlerCustomer struct {
	customer entity.EntityCustomer
}

func NewHandlerCustomer(customer entity.EntityCustomer) *handlerCustomer {
	return &handlerCustomer{customer: customer}
}

func (h *handlerCustomer) HandlerHello(c *fiber.Ctx) error {
	return c.SendString("Handler customer")

}

func (h *handlerCustomer) HandlerCreate(c *fiber.Ctx) error {
	var body schemas.SchemaCustomer
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.customer.EntityCreate(&body)

	if error.Type == "error_create_01" {
		return c.Status(error.Code).JSON(fiber.Map{
			"error": error.Code,
			"msg":   "Error name already",
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "create successfully",
	})
}

func (h *handlerCustomer) HandlerResults(c *fiber.Ctx) error {
	res, err := h.customer.EntityResults()

	if err.Type == "error_results_01" {
		return c.Status(err.Code).JSON(fiber.Map{
			"msg":   "Customer Not found",
			"error": err.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Customer data already to use",
		"status": true,
		"data":   res,
	})
}

func (h *handlerCustomer) HandlerResult(c *fiber.Ctx) error {
	var body schemas.SchemaCustomer

	id := c.Params("id")
	body.ID = id

	res, error := h.customer.EntityResult(&body)

	if error.Type == "error_result_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "Customer not found",
			"error": error.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Customer found",
		"status": true,
		"data":   res,
	})
}

func (h *handlerCustomer) HandlerDelete(c *fiber.Ctx) error {
	var body schemas.SchemaCustomer

	id := c.Params("id")
	body.ID = id

	res, error := h.customer.EntityResult(&body)

	if error.Type == "error_delete_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "customer not found",
			"error": error.Code,
		})
	}

	if error.Type == "error_delete_02" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "customer id failed",
			"error": error.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Category  found",
		"status": true,
		"data":   res,
	})
}

func (h *handlerCustomer) HandlerUpdate(c *fiber.Ctx) error {
	var body schemas.SchemaCustomer
	id := c.Params("id")

	body.ID = id

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.customer.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "customer not found",
			"error": error.Code,
		})
	}

	if error.Type == "error_delete_02" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "customer category id failed",
			"error": error.Code,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Update customer data success for this id",
		"code": fiber.StatusCreated,
	})
}

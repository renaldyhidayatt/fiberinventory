package handler

import (
	"fiberinventory/entity"
	"fiberinventory/schemas"

	"github.com/gofiber/fiber/v2"
)

type handlerSupplier struct {
	supplier entity.EntitySupplier
}

func NewHandlerSupplier(supplier entity.EntitySupplier) *handlerSupplier {
	return &handlerSupplier{supplier: supplier}
}

func (h *handlerSupplier) HandlerHello(c *fiber.Ctx) error {
	return c.SendString("Handler Supplier")

}

func (h *handlerSupplier) HandlerCreate(c *fiber.Ctx) error {
	var body schemas.SchemaSupplier
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.supplier.EntityCreate(&body)

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

func (h *handlerSupplier) HandlerResults(c *fiber.Ctx) error {
	res, err := h.supplier.EntityResults()

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

func (h *handlerSupplier) HandlerResult(c *fiber.Ctx) error {
	var body schemas.SchemaSupplier

	id := c.Params("id")
	body.ID = id

	res, error := h.supplier.EntityResult(&body)

	if error.Type == "error_result_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "Supplier not found",
			"error": error.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Customer found",
		"status": true,
		"data":   res,
	})
}

func (h *handlerSupplier) HandlerDelete(c *fiber.Ctx) error {
	var body schemas.SchemaSupplier

	id := c.Params("id")
	body.ID = id

	res, error := h.supplier.EntityResult(&body)

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

func (h *handlerSupplier) HandlerUpdate(c *fiber.Ctx) error {
	var body schemas.SchemaSupplier
	id := c.Params("id")

	body.ID = id

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.supplier.EntityUpdate(&body)

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

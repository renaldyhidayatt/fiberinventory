package handler

import (
	"fiberinventory/entity"
	"fiberinventory/schemas"

	"github.com/gofiber/fiber/v2"
)

type handlerProduct struct {
	product entity.EntityProduct
}

func NewHandlerProduct(product entity.EntityProduct) *handlerProduct {
	return &handlerProduct{product: product}
}

func (h *handlerProduct) HandlerHello(c *fiber.Ctx) error {
	return c.SendString("Handler Product")
}

func (h *handlerProduct) HandlerCreate(c *fiber.Ctx) error {
	var body schemas.SchemaProduct

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.product.EntityCreate(&body)

	if error.Type == "error_create_01" {
		return c.Status(error.Code).JSON(fiber.Map{
			"error": error.Code,
			"msg":   "Error Name Already",
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "create sucessfully",
	})
}

func (h *handlerProduct) HandlerResults(c *fiber.Ctx) error {
	res, err := h.product.EntityResults()

	if err.Type == "error_results_01" {
		return c.Status(err.Code).JSON(fiber.Map{
			"msg":   "Product Not Found",
			"error": err.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Product data already to use",
		"status": true,
		"data":   res,
	})
}

func (h *handlerProduct) HandlerResult(c *fiber.Ctx) error {
	var body schemas.SchemaProduct

	id := c.Params("id")

	body.ID = id

	res, error := h.product.EntityResult(&body)

	if error.Type == "error_result_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "Product Not found",
			"error": error.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Product found",
		"status": true,
		"data":   res,
	})
}

func (h *handlerProduct) HandlerDelete(c *fiber.Ctx) error {
	var body schemas.SchemaProduct

	id := c.Params("id")

	body.ID = id

	res, error := h.product.EntityDelete(&body)

	if error.Type == "error_delete_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "Product Not found",
			"error": error.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Customer found",
		"status": true,
		"data":   res,
	})
}

func (h *handlerProduct) HandlerUpdate(c *fiber.Ctx) error {
	var body schemas.SchemaProduct

	id := c.Params("id")

	body.ID = id

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.product.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "product not found",
			"error": error.Code,
		})
	}

	if error.Type == "error_delete_02" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "product id failed",
			"error": error.Code,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Update product data success for this id",
		"code": fiber.StatusCreated,
	})

}

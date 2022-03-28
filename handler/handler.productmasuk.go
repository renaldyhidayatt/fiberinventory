package handler

import (
	"fiberinventory/entity"
	"fiberinventory/schemas"

	"github.com/gofiber/fiber/v2"
)

type handlerProductMasuk struct {
	productmasuk entity.EntityProductMasuk
}

func NewHandlerProductMasuk(productmasuk entity.EntityProductMasuk) *handlerProductMasuk {
	return &handlerProductMasuk{productmasuk: productmasuk}
}

func (h *handlerProductMasuk) HandlerHello(c *fiber.Ctx) error {
	return c.SendString("Handler ProductMasuk")
}

func (h *handlerProductMasuk) HandlerCreate(c *fiber.Ctx) error {
	var body schemas.SchemaProductMasuk

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.productmasuk.EntityCreate(&body)

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

func (h *handlerProductMasuk) HandlerResults(c *fiber.Ctx) error {
	res, err := h.productmasuk.EntityResults()

	if err.Type == "error_results_01" {
		return c.Status(err.Code).JSON(fiber.Map{
			"msg":   "ProductMasuk Not found",
			"error": err.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "productmasuk data ready",
		"status": true,
		"data":   res,
	})
}

func (h *handlerProductMasuk) HandlerResult(c *fiber.Ctx) error {
	var body schemas.SchemaProductMasuk

	id := c.Params("id")

	body.ID = id

	res, error := h.productmasuk.EntityResult(&body)

	if error.Type == "error_result_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "ProductMasuk Not found",
			"error": error.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "ProductMasuk found",
		"status": true,
		"data":   res,
	})
}

func (h *handlerProductMasuk) HandlerDelete(c *fiber.Ctx) error {
	var body schemas.SchemaProductMasuk

	id := c.Params("id")
	body.ID = id

	res, error := h.productmasuk.EntityResult(&body)

	if error.Type == "error_delete_01" {
		return c.Status(error.Code).JSON(fiber.Map{
			"msg":   "product not found",
			"error": error.Code,
		})
	}

	if error.Type == "error_delete_02" {
		return c.Status(error.Code).JSON(fiber.Map{
			"msg":   "product id failed",
			"error": error.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Product found",
		"status": true,
		"data":   res,
	})
}

func (h *handlerProductMasuk) HandlerUpdate(c *fiber.Ctx) error {
	var body schemas.SchemaProductMasuk
	id := c.Params("id")

	body.ID = id

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.productmasuk.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "product not found",
			"error": error.Code,
		})
	}

	if error.Type == "error_delete_02" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "product category id failed",
			"error": error.Code,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Update product data success for this id",
		"code": fiber.StatusCreated,
	})
}

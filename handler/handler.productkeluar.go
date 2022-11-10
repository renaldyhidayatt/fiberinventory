package handler

import (
	"fiberinventory/entity"
	"fiberinventory/schemas"

	"github.com/gofiber/fiber/v2"
)

type handlerProductKeluar struct {
	productkeluar entity.EntityProductKeluar
}

func NewHandlerProductKeluar(productkeluar entity.EntityProductKeluar) *handlerProductKeluar {
	return &handlerProductKeluar{productkeluar: productkeluar}
}

func (h *handlerProductKeluar) HandlerCreate(c *fiber.Ctx) error {
	var body schemas.SchemaProductKeluar

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.productkeluar.EntityCreate(&body)

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

func (h *handlerProductKeluar) HandlerResults(c *fiber.Ctx) error {
	res, err := h.productkeluar.EntityResults()

	if err.Type == "error_results_01" {
		return c.Status(err.Code).JSON(fiber.Map{
			"msg":   "ProductKeluar Not found",
			"error": err.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "productkeluar data ready",
		"status": true,
		"data":   res,
	})
}

func (h *handlerProductKeluar) HandlerResult(c *fiber.Ctx) error {
	var body schemas.SchemaProductKeluar

	id := c.Params("id")

	body.ID = id

	res, error := h.productkeluar.EntityResult(&body)

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

func (h *handlerProductKeluar) HandlerDelete(c *fiber.Ctx) error {
	var body schemas.SchemaProductKeluar

	id := c.Params("id")
	body.ID = id

	res, error := h.productkeluar.EntityResult(&body)

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

func (h *handlerProductKeluar) HandlerUpdate(c *fiber.Ctx) error {
	var body schemas.SchemaProductKeluar
	id := c.Params("id")

	body.ID = id

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.productkeluar.EntityUpdate(&body)

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

package handler

import (
	"fiberinventory/entity"
	"fiberinventory/schemas"

	"github.com/gofiber/fiber/v2"
)

type handlerSale struct {
	sale entity.EntitySale
}

func NewHandlerSale(sale entity.EntitySale) *handlerSale {
	return &handlerSale{sale: sale}
}

func (h *handlerSale) HandlerHello(c *fiber.Ctx) error {

	return c.SendString("Handler Hello")
}

func (h *handlerSale) HandlerCreate(c *fiber.Ctx) error {
	var body schemas.SchemaSale
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.sale.EntityCreate(&body)

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

func (h *handlerSale) HandlerResults(c *fiber.Ctx) error {
	res, err := h.sale.EntityResults()

	if err.Type == "error_results_01" {
		return c.Status(err.Code).JSON(fiber.Map{
			"msg":   "Sale Not found",
			"error": err.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Sale data already to use",
		"status": true,
		"data":   res,
	})
}

func (h *handlerSale) HandlerResult(c *fiber.Ctx) error {
	var body schemas.SchemaSale

	id := c.Params("id")
	body.ID = id

	res, error := h.sale.EntityResult(&body)

	if error.Type == "error_result_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "Sale not found",
			"error": error.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "sale found",
		"status": true,
		"data":   res,
	})
}

func (h *handlerSale) HandlerDelete(c *fiber.Ctx) error {
	var body schemas.SchemaSale

	id := c.Params("id")
	body.ID = id

	res, error := h.sale.EntityResult(&body)

	if error.Type == "error_delete_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "sale not found",
			"error": error.Code,
		})
	}

	if error.Type == "error_delete_02" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "sale id failed",
			"error": error.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "sale  found",
		"status": true,
		"data":   res,
	})
}

func (h *handlerSale) HandlerUpdate(c *fiber.Ctx) error {
	var body schemas.SchemaSale
	id := c.Params("id")

	body.ID = id

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.sale.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "sale not found",
			"error": error.Code,
		})
	}

	if error.Type == "error_delete_02" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "sale id failed",
			"error": error.Code,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Update customer data success for this id",
		"code": fiber.StatusCreated,
	})
}

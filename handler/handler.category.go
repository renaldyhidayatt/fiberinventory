package handler

import (
	"fiberinventory/entity"
	"fiberinventory/schemas"

	"github.com/gofiber/fiber/v2"
)

type handlerCategory struct {
	category entity.EntityCategory
}

func NewHandlerCategory(category entity.EntityCategory) *handlerCategory {
	return &handlerCategory{category: category}
}

func (h *handlerCategory) HandlerHello(c *fiber.Ctx) error {
	return c.SendString("Handler merchant")
}

func (h *handlerCategory) HandlerCreate(c *fiber.Ctx) error {
	var body schemas.SchemaCategory
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.category.EntityCreate(&body)

	if error.Type == "error_create_01" {
		return c.Status(error.Code).JSON(fiber.Map{
			"error": error.Code,
			"msg":   "Error name already",
		})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      "create successfuly",
		"merchant": body,
	})
}

func (h *handlerCategory) HandlerResults(c *fiber.Ctx) error {
	res, err := h.category.EntityResults()

	if err.Type == "error_results_01" {
		return c.Status(err.Code).JSON(fiber.Map{
			"msg":   "Category Not found",
			"error": err.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Category data already to use",
		"status": true,
		"data":   res,
	})
}

func (h *handlerCategory) HandlerResult(c *fiber.Ctx) error {
	var body schemas.SchemaCategory

	id := c.Params("id")
	body.ID = id

	res, error := h.category.EntityResult(&body)

	if error.Type == "error_result_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "category not found",
			"error": error.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Category  found",
		"status": true,
		"data":   res,
	})
}

func (h *handlerCategory) HandlerDelete(c *fiber.Ctx) error {
	var body schemas.SchemaCategory

	id := c.Params("id")
	body.ID = id

	res, error := h.category.EntityResult(&body)

	if error.Type == "error_delete_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "category not found",
			"error": error.Code,
		})
	}

	if error.Type == "error_delete_02" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "category category id failed",
			"error": error.Code,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Category not found",
		"status": true,
		"data":   res,
	})
}

func (h *handlerCategory) HandlerUpdate(c *fiber.Ctx) error {
	var body schemas.SchemaCategory
	id := c.Params("id")

	body.ID = id

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	_, error := h.category.EntityUpdate(&body)

	if error.Type == "error_update_01" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "category not found",
			"error": error.Code,
		})
	}

	if error.Type == "error_delete_02" {
		c.Status(error.Code).JSON(fiber.Map{
			"msg":   "category category id failed",
			"error": error.Code,
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Update customer data success for this id",
		"code": fiber.StatusCreated,
	})
}

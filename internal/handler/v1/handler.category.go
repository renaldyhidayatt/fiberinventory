package v1

import (
	"fiberinventory/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initCategoryGroup(api *fiber.App) {
	routerCategory := api.Group("/category")

	routerCategory.Get("/hello", h.handlerCategory)
	routerCategory.Post("/create", h.handlerCreate)
	routerCategory.Put("/update/:id", h.HandlerUpdate)
	routerCategory.Delete("/delete/:id", h.HandlerDelete)
	routerCategory.Get("/", h.handleresults)
	routerCategory.Get("/:id", h.HandlerResult)
}

func (h *Handler) handlerCategory(c *fiber.Ctx) error {
	return c.SendString("Handler merchant")
}

func (h *Handler) handlerCreate(c *fiber.Ctx) error {
	var body domain.CategoryInput
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

	res, err := h.services.Category.Create(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error":    false,
		"msg":      "create successfuly",
		"merchant": res,
	})
}

func (h *Handler) handleresults(c *fiber.Ctx) error {
	res, err := h.services.Category.Results()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":   err.Error(),
			"error": true,
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Category data already to use",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) HandlerResult(c *fiber.Ctx) error {
	var body domain.CategoryInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Category.Result(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Category  found",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) HandlerDelete(c *fiber.Ctx) error {
	var body domain.CategoryInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Category.Delete(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error":   true,
			"message": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Success delete category",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) HandlerUpdate(c *fiber.Ctx) error {
	var body domain.CategoryInput
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

	res, err := h.services.Category.Update(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
		"msg":  "Update category data success for this id",
		"code": fiber.StatusAccepted,
	})
}

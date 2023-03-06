package v1

import (
	"fiberinventory/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initProductMasukGroup(api *fiber.App) {
	routerProductMasuk := api.Group("/productmasuk")

	routerProductMasuk.Get("/", h.handlerProductMasukResults)
	routerProductMasuk.Get("/:id", h.handlerProductMasukResult)
	routerProductMasuk.Post("/create", h.handlerProductMasukCreate)
	routerProductMasuk.Post("/update/:id", h.handlerProductMasukUpdate)
	routerProductMasuk.Post("/delete/:id", h.handlerProductMasukDelete)

}

func (h *Handler) handlerProductMasukCreate(c *fiber.Ctx) error {
	var body domain.ProductMasukInput

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

	res, err := h.services.ProductMasuk.Create(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"data":  res,
		"error": false,
		"msg":   "create successfully",
	})
}

func (h *Handler) handlerProductMasukResults(c *fiber.Ctx) error {
	res, err := h.services.ProductMasuk.Results()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "productmasuk data ready",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerProductMasukResult(c *fiber.Ctx) error {
	var body domain.ProductMasukInput

	id := c.Params("id")

	body.ID = id

	res, err := h.services.ProductMasuk.Result(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "ProductMasuk found",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerProductMasukDelete(c *fiber.Ctx) error {
	var body domain.ProductMasukInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.ProductMasuk.Delete(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Product found",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerProductMasukUpdate(c *fiber.Ctx) error {
	var body domain.ProductMasukInput
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
	res, err := h.services.ProductMasuk.Update(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "Update product data success for this id",
		"code": fiber.StatusCreated,
		"data": res,
	})
}

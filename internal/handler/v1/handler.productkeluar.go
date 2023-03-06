package v1

import (
	"fiberinventory/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initProductKeluarGroup(api *fiber.App) {
	routerProductKeluar := api.Group("/productkeluar")

	routerProductKeluar.Get("/", h.handlerProductKeluarResults)
	routerProductKeluar.Get("/:id", h.handlerProductKeluarResult)
	routerProductKeluar.Post("/create", h.handlerProductKeluarCreate)
	routerProductKeluar.Post("/update/:id", h.handlerProductKeluarUpdate)
	routerProductKeluar.Post("/delete/:id", h.handlerProductKeluarDelete)

}

func (h *Handler) handlerProductKeluarCreate(c *fiber.Ctx) error {
	var body domain.ProductKeluarInput

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

	res, err := h.services.ProductKeluar.Create(&body)

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

func (h *Handler) handlerProductKeluarResults(c *fiber.Ctx) error {
	res, err := h.services.ProductKeluar.Results()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "productkeluar data ready",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerProductKeluarResult(c *fiber.Ctx) error {
	var body domain.ProductKeluarInput

	id := c.Params("id")

	body.ID = id

	res, err := h.services.ProductKeluar.Result(&body)

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

func (h *Handler) handlerProductKeluarDelete(c *fiber.Ctx) error {
	var body domain.ProductKeluarInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.ProductKeluar.Delete(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Success delete product keluar",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerProductKeluarUpdate(c *fiber.Ctx) error {
	var body domain.ProductKeluarInput
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

	res, err := h.services.ProductKeluar.Update(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
		"msg":  "Update product data success for this id",
		"code": fiber.StatusCreated,
	})
}

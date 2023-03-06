package v1

import (
	"fiberinventory/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initSupplierGroup(api *fiber.App) {
	routerSupplier := api.Group("/supplier")

	routerSupplier.Get("/", h.handlerSupplierResults)
	routerSupplier.Get("/:id", h.handlerSupplierResult)
	routerSupplier.Post("/create", h.handlerSupplierCreate)
	routerSupplier.Post("/update/:id", h.handlerSupplierUpdate)
	routerSupplier.Post("/delete/:id", h.handlerSupplierDelete)

}

func (h *Handler) handlerSupplierCreate(c *fiber.Ctx) error {
	var body domain.SupplierInput

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

	res, err := h.services.Supplier.Create(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"error": false,
		"msg":   "create successfully",
		"data":  res,
	})
}

func (h *Handler) handlerSupplierResults(c *fiber.Ctx) error {
	res, err := h.services.Supplier.Results()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Customer data already to use",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerSupplierResult(c *fiber.Ctx) error {
	var body domain.SupplierInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Supplier.Result(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Customer found",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerSupplierDelete(c *fiber.Ctx) error {
	var body domain.SupplierInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Supplier.Delete(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Supplier found",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerSupplierUpdate(c *fiber.Ctx) error {
	var body domain.SupplierInput

	id := c.Params("id")

	body.ID = id

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	res, err := h.services.Supplier.Update(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
		"msg":  "Update customer data success for this id",
		"code": fiber.StatusCreated,
	})
}

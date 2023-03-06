package v1

import (
	"fiberinventory/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initSaleGroup(api *fiber.App) {
	routerSale := api.Group("/sale")

	routerSale.Get("/", h.handlerSaleResults)
	routerSale.Get("/:id", h.handlerSaleResult)
	routerSale.Post("/create", h.handlerSaleCreate)
	routerSale.Post("/update/:id", h.handlerSaleUpdate)
	routerSale.Post("/delete/:id", h.handlerSaleDelete)

}

func (h *Handler) handlerSaleCreate(c *fiber.Ctx) error {
	var body domain.SaleInput

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

	res, err := h.services.Sale.Create(&body)

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

func (h *Handler) handlerSaleResults(c *fiber.Ctx) error {
	res, err := h.services.Sale.Results()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Sale data already to use",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerSaleResult(c *fiber.Ctx) error {
	var body domain.SaleInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Sale.Result(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "sale found",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerSaleDelete(c *fiber.Ctx) error {
	var body domain.SaleInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Sale.Delete(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "sale  found",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerSaleUpdate(c *fiber.Ctx) error {
	var body domain.SaleInput
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

	res, err := h.services.Sale.Update(&body)

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

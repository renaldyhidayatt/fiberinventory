package v1

import (
	"fiberinventory/internal/domain"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initProductGroup(api *fiber.App) {
	routerProduct := api.Group("/product")

	routerProduct.Get("/hello", h.handlerProductHello)
	routerProduct.Get("/", h.handlerProductResults)
	routerProduct.Get("/:id", h.handlerProductResult)
	routerProduct.Post("/create", h.handlerProductCreate)
	routerProduct.Post("/update/:id", h.handlerProductUpdate)
	routerProduct.Post("/delete/:id", h.handlerProductDelete)

}

func (h *Handler) handlerProductHello(c *fiber.Ctx) error {
	return c.SendString("Handler Product")
}

func (h *Handler) handlerProductCreate(c *fiber.Ctx) error {
	var body domain.ProductInput

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

	res, err := h.services.Product.Create(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":   err.Error(),
			"error": true,
		})
	}

	return c.JSON(fiber.Map{
		"data":  res,
		"error": false,
		"msg":   "create sucessfully",
	})
}

func (h *Handler) handlerProductResults(c *fiber.Ctx) error {
	res, err := h.services.Product.Results()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Product data already to use",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerProductResult(c *fiber.Ctx) error {
	var body domain.ProductInput

	id := c.Params("id")

	body.ID = id

	res, err := h.services.Product.Result(&body)

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

func (h *Handler) handlerProductDelete(c *fiber.Ctx) error {
	var body domain.ProductInput

	id := c.Params("id")

	body.ID = id

	res, err := h.services.Product.Delete(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": true,
			"msg":   err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"msg":    "Success delete Product",
		"status": true,
		"data":   res,
	})
}

func (h *Handler) handlerProductUpdate(c *fiber.Ctx) error {
	var body domain.ProductInput

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

	res, err := h.services.Product.Update(&body)

	if err != nil {
		if err := body.Validate(); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"data": res,
		"msg":  "Update product data success for this id",
		"code": fiber.StatusCreated,
	})

}

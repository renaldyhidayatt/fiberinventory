package v1

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initProductGroup(api *fiber.App) {
	routerProduct := api.Group("/product")

	routerProduct.Use(middleware.Proctected())

	routerProduct.Get("/hello", h.handlerProductHello)
	routerProduct.Get("/", h.handlerProductResults)
	routerProduct.Get("/:id", h.handlerProductResult)
	routerProduct.Post("/create", h.handlerProductCreate)
	routerProduct.Post("/update/:id", h.handlerProductUpdate)
	routerProduct.Post("/delete/:id", h.handlerProductDelete)

}

// handlerProductHello function
// @Summary Greet the Product
// @Description Return a greeting message to the product
// @Tags Product
// @Produce plain
// @Security BearerAuth
// @Success 200 {string} string "OK"
// @Router /product/hello [get]
func (h *Handler) handlerProductHello(c *fiber.Ctx) error {
	return c.SendString("Handler Product")
}

// handlerProductCreate function
// @Summary createProduct to the application
// @Description Create Product
// @Tags product
// @Accept json
// @Produce json
// @Param product body domain.ProductInput true "Product information"
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /product/create [post]
func (h *Handler) handlerProductCreate(c *fiber.Ctx) error {
	var body domain.ProductInput

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Product.Create(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil membuat category",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// handlerProductResults function
// @Summary Get product results
// @Description Retrieve the results for each product
// @Tags Product
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /product [get]
func (h *Handler) handlerProductResults(c *fiber.Ctx) error {
	res, err := h.services.Product.Results()

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil membuat category",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// handlerProductResult function
// @Summary Get product result
// @Description Retrieve the result for a specific product
// @Tags Product
// @Param id path string true "Product ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /product/{id} [get]
func (h *Handler) handlerProductResult(c *fiber.Ctx) error {
	var body domain.ProductInput

	id := c.Params("id")

	body.ID = id

	res, err := h.services.Product.Result(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil mendapatkan data",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// handlerProductDelete function
// @Summary Delete product
// @Description Delete a specific product
// @Tags Product
// @Param id path string true "Product ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /product/{id} [delete]
func (h *Handler) handlerProductDelete(c *fiber.Ctx) error {
	var body domain.ProductInput

	id := c.Params("id")

	body.ID = id

	res, err := h.services.Product.Delete(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil menghapus category",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// handlerProductUpdate function
// @Summary Update product
// @Description Update a specific product
// @Tags Product
// @Param id path string true "Product ID"
// @Param body body domain.ProductInput true "Product Data"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /product/{id} [put]
func (h *Handler) handlerProductUpdate(c *fiber.Ctx) error {
	var body domain.ProductInput

	id := c.Params("id")

	body.ID = id

	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}
	if err := body.Validate(); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	res, err := h.services.Product.Update(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil mengupdate category",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})

}

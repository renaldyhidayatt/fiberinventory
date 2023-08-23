package api

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"
	"fiberinventory/internal/pb"

	"github.com/gofiber/fiber/v2"
)

type handleProduct struct {
	client pb.ProductServiceClient
}

func NewHandleProduct(client pb.ProductServiceClient, router *fiber.App) {

	handlerProduct := handleProduct{
		client: client,
	}

	routerProduct := router.Group("/api/product")

	routerProduct.Use(middleware.Proctected())

	routerProduct.Get("/hello", handlerProduct.handlerProductHello)
	routerProduct.Get("/", handlerProduct.handlerProductResults)
	routerProduct.Get("/:id", handlerProduct.handlerProductResult)
	routerProduct.Post("/create", handlerProduct.handlerProductCreate)
	routerProduct.Post("/update/:id", handlerProduct.handlerProductUpdate)
	routerProduct.Post("/delete/:id", handlerProduct.handlerProductDelete)

}

// handlerProductHello function
// @Summary Greet the Product
// @Description Return a greeting message to the product
// @Tags Product
// @Produce plain
// @Security BearerAuth
// @Success 200 {string} string "OK"
// @Router /product/hello [get]
func (h *handleProduct) handlerProductHello(c *fiber.Ctx) error {
	return c.SendString("Handler Product")
}

// handlerProductCreate function
// @Summary createProduct to the application
// @Description Create Product
// @Tags product
// @Accept json
// @Produce json
// @Param product body domain.CreateProductRequest true "Product information"
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /product/create [post]
func (h *handleProduct) handlerProductCreate(c *fiber.Ctx) error {
	var body domain.CreateProductRequest

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

	data := &pb.CreateProductInput{
		Name:       body.Name,
		Image:      body.Image,
		Qty:        body.Qty,
		CategoryId: body.CategoryID,
	}

	res, err := h.client.CreateProduct(c.Context(), data)

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
func (h *handleProduct) handlerProductResults(c *fiber.Ctx) error {
	res, err := h.client.GetProducts(c.Context(), &pb.ProductsRequest{})

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
func (h *handleProduct) handlerProductResult(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.ProductRequest{
		Id: id,
	}

	res, err := h.client.GetProduct(c.Context(), data)

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
func (h *handleProduct) handlerProductDelete(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.ProductRequest{
		Id: id,
	}

	res, err := h.client.DeleteProduct(c.Context(), data)

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
// @Param body body domain.UpdateProductRequest true "Product Data"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /product/{id} [put]
func (h *handleProduct) handlerProductUpdate(c *fiber.Ctx) error {
	var body domain.UpdateProductRequest

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

	data := &pb.UpdateProductInput{
		Id:         body.ID,
		Name:       body.Name,
		Image:      body.Image,
		Qty:        body.Qty,
		CategoryId: body.CategoryID,
	}

	res, err := h.client.UpdateProduct(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil mengupdate product",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})

}

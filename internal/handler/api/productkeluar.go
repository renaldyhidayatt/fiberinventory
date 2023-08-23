package api

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"
	"fiberinventory/internal/pb"

	"github.com/gofiber/fiber/v2"
)

type handleProductKeluar struct {
	client pb.ProductKeluarServiceClient
}

func NewHandleProductKeluar(client pb.ProductKeluarServiceClient, router *fiber.App) {
	h := handleProductKeluar{
		client: client,
	}

	routerProductKeluar := router.Group("/api/productkeluar")

	routerProductKeluar.Use(middleware.Proctected())

	routerProductKeluar.Get("/hello", h.handlerProductKeluarHello)
	routerProductKeluar.Get("/", h.handlerProductKeluarResults)
	routerProductKeluar.Get("/:id", h.handlerProductKeluarResult)
	routerProductKeluar.Post("/create", h.handlerProductKeluarCreate)
	routerProductKeluar.Post("/update/:id", h.handlerProductKeluarUpdate)
	routerProductKeluar.Post("/delete/:id", h.handlerProductKeluarDelete)
}

// handlerProductKeluarHello function
// @Summary Greet the ProductKeluar
// @Description Return a greeting message for the ProductKeluar
// @Tags ProductKeluar
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /productkeluar/hello [get]
func (h *handleProductKeluar) handlerProductKeluarHello(c *fiber.Ctx) error {
	return c.SendString("Handler ProductKeluar")
}

// handlerProductKeluarCreate function
// @Summary Create a new ProductKeluar
// @Description Create ProductKeluar
// @Tags ProductKeluar
// @Accept json
// @Produce json
// @Param productkeluar body domain.CreateProductKeluarRequest true "ProductKeluar information"
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productkeluar/create [post]
func (h *handleProductKeluar) handlerProductKeluarCreate(c *fiber.Ctx) error {
	var body domain.CreateProductKeluarRequest

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

	data := &pb.CreateProductKeluarInput{
		Qty:        body.Qty,
		ProductId:  body.ProductID,
		CategoryId: body.CategoryID,
	}

	res, err := h.client.CreateProductKeluar(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil membuat product keluar",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// handlerProductKeluarResults function
// @Summary Get ProductKeluar results
// @Description Retrieve the results for all ProductKeluar
// @Tags ProductKeluar
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productkeluar [get]
func (h *handleProductKeluar) handlerProductKeluarResults(c *fiber.Ctx) error {
	res, err := h.client.GetProductKeluars(c.Context(), &pb.ProductKeluarsRequest{})

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

// handlerProductKeluarResult function
// @Summary Get ProductKeluar result
// @Description Retrieve the result for a specific ProductKeluar
// @Tags ProductKeluar
// @Param id path string true "ProductKeluar ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productkeluar/{id} [get]
func (h *handleProductKeluar) handlerProductKeluarResult(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.ProductKeluarRequest{
		Id: id,
	}

	res, err := h.client.GetProductKeluar(c.Context(), data)

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

// handlerProductKeluarDelete function
// @Summary Delete a ProductKeluar
// @Description Delete a specific ProductKeluar
// @Tags ProductKeluar
// @Param id path string true "ProductKeluar ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productkeluar/{id} [delete]
func (h *handleProductKeluar) handlerProductKeluarDelete(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.ProductKeluarRequest{
		Id: id,
	}

	res, err := h.client.DeleteProductKeluar(c.Context(), data)

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

// handlerProductKeluarUpdate function
// @Summary Update a ProductKeluar
// @Description Update a specific ProductKeluar
// @Tags ProductKeluar
// @Param id path string true "ProductKeluar ID"
// @Param body body domain.UpdateProductKeluarRequest true "ProductKeluar Data"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productkeluar/{id} [put]
func (h *handleProductKeluar) handlerProductKeluarUpdate(c *fiber.Ctx) error {
	var body domain.UpdateProductKeluarRequest
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

	data := &pb.UpdateProductKeluarInput{
		Id:         body.ID,
		Qty:        body.Qty,
		ProductId:  body.ProductID,
		CategoryId: body.CategoryID,
	}

	res, err := h.client.UpdateProductKeluar(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil mengupdate product keluar",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

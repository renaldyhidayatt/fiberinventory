package api

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"
	"fiberinventory/internal/pb"

	"github.com/gofiber/fiber/v2"
)

type handleProductMasuk struct {
	client pb.ProductMasukServiceClient
}

func NewHandleProductMasuk(client pb.ProductMasukServiceClient, router *fiber.App) {
	handlerProductMasuk := handleProductMasuk{
		client: client,
	}

	routerProductMasuk := router.Group("/api/productmasuk")

	routerProductMasuk.Use(middleware.Proctected())

	routerProductMasuk.Get("/hello", handlerProductMasuk.handlerProductMasukHello)
	routerProductMasuk.Get("/", handlerProductMasuk.handlerProductMasukResults)
	routerProductMasuk.Get("/:id", handlerProductMasuk.handlerProductMasukResult)
	routerProductMasuk.Post("/create", handlerProductMasuk.handlerProductMasukCreate)
	routerProductMasuk.Post("/update/:id", handlerProductMasuk.handlerProductMasukUpdate)
	routerProductMasuk.Post("/delete/:id", handlerProductMasuk.handlerProductMasukDelete)
}

// handlerProductMasukHello function
// @Summary Greet the ProductMasuk
// @Description Return a greeting message for the ProductMasuk
// @Tags ProductMasuk
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /productmasuk/hello [get]
func (h *handleProductMasuk) handlerProductMasukHello(c *fiber.Ctx) error {
	return c.SendString("Handler ProductMasuk")
}

// handlerProductMasukCreate function
// @Summary Create a new ProductMasuk
// @Description Create ProductMasuk
// @Tags ProductMasuk
// @Accept json
// @Produce json
// @Param productmasuk body domain.CreateProductMasukRequest true "ProductMasuk information"
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productmasuk/create [post]
func (h *handleProductMasuk) handlerProductMasukCreate(c *fiber.Ctx) error {
	var body domain.CreateProductMasukRequest

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

	data := &pb.CreateProductMasukInput{
		Name:       body.Name,
		Qty:        body.Qty,
		ProductId:  body.ProductID,
		SupplierId: body.SupplierID,
	}

	res, err := h.client.CreateProductMasuk(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil membuat product masuk",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// handlerProductMasukResults function
// @Summary Get ProductMasuk results
// @Description Retrieve the results for all ProductMasuk
// @Tags ProductMasuk
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productmasuk [get]
func (h *handleProductMasuk) handlerProductMasukResults(c *fiber.Ctx) error {
	res, err := h.client.GetProductMasuks(c.Context(), &pb.ProductMasuksRequest{})

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

// handlerProductMasukResult function
// @Summary Get ProductMasuk result
// @Description Retrieve the result for a specific ProductMasuk
// @Tags ProductMasuk
// @Param id path string true "ProductMasuk ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /api/productmasuk/{id} [get]
func (h *handleProductMasuk) handlerProductMasukResult(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.ProductMasukRequest{
		Id: id,
	}

	res, err := h.client.GetProductMasuk(c.Context(), data)

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

// handlerProductMasukDelete function
// @Summary Delete ProductMasuk
// @Description Delete a specific ProductMasuk
// @Tags ProductMasuk
// @Param id path string true "ProductMasuk ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /api/productmasuk/{id} [delete]
func (h *handleProductMasuk) handlerProductMasukDelete(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.ProductMasukRequest{
		Id: id,
	}
	res, err := h.client.DeleteProductMasuk(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil menghapus productmasuk",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// handlerProductMasukUpdate function
// @Summary Update ProductMasuk
// @Description Update a specific ProductMasuk
// @Tags ProductMasuk
// @Param id path string true "ProductMasuk ID"
// @Param body body domain.UpdateProductMasukRequest true "ProductMasuk Data"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /api/productmasuk/{id} [put]
func (h *handleProductMasuk) handlerProductMasukUpdate(c *fiber.Ctx) error {
	var body domain.UpdateProductMasukRequest
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
	data := &pb.UpdateProductMasukInput{
		Id:         body.ID,
		Name:       body.Name,
		Qty:        body.Qty,
		ProductId:  body.ProductID,
		SupplierId: body.SupplierID,
	}

	res, err := h.client.UpdateProductMasuk(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil mengupdate productmasuk",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

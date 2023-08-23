package api

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"
	"fiberinventory/internal/pb"

	"github.com/gofiber/fiber/v2"
)

type handleSale struct {
	client pb.SaleServiceClient
}

func NewHandlerSale(client pb.SaleServiceClient, router *fiber.App) {
	h := handleSale{
		client: client,
	}

	routerSale := router.Group("/api/sale")

	routerSale.Use(middleware.Proctected())

	routerSale.Get("/hello", h.handlerSaleHello)
	routerSale.Get("/", h.handlerSaleResults)
	routerSale.Get("/:id", h.handlerSaleResult)
	routerSale.Post("/create", h.handlerSaleCreate)
	routerSale.Post("/update/:id", h.handlerSaleUpdate)
	routerSale.Post("/delete/:id", h.handlerSaleDelete)

}

// handlerSaleHello function
// @Summary Greet the Sale
// @Description Return a greeting message for the Sale
// @Tags Sale
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /sale/hello [get]
func (h *handleSale) handlerSaleHello(c *fiber.Ctx) error {
	return c.SendString("Handler Sale")
}

// handlerSaleCreate function
// @Summary Create a Sale
// @Description Create a new Sale
// @Tags Sale
// @Accept json
// @Produce json
// @Param body body domain.CreateSaleRequest true "Sale information"
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /sale/create [post]
func (h *handleSale) handlerSaleCreate(c *fiber.Ctx) error {
	var body domain.CreateSaleRequest

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

	data := &pb.CreateSaleInput{
		Name:    body.Name,
		Alamat:  body.Alamat,
		Email:   body.Email,
		Telepon: body.Telepon,
	}

	res, err := h.client.CreateSale(c.Context(), data)

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

// handlerSaleResults function
// @Summary Get Sale results
// @Description Retrieve the results for each Sale
// @Tags Sale
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /sale [get]
func (h *handleSale) handlerSaleResults(c *fiber.Ctx) error {
	res, err := h.client.GetSales(c.Context(), &pb.SalesRequest{})

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

// handlerSaleResult function
// @Summary Get Sale result
// @Description Retrieve the result for a specific Sale
// @Tags Sale
// @Param id path string true "Sale ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /sale/{id} [get]
func (h *handleSale) handlerSaleResult(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.SaleRequest{
		Id: id,
	}

	res, err := h.client.GetSale(c.Context(), data)

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

// handlerSaleDelete function
// @Summary Delete Sale
// @Description Delete a specific Sale
// @Tags Sale
// @Param id path string true "Sale ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /sale/{id} [delete]
func (h *handleSale) handlerSaleDelete(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.SaleRequest{
		Id: id,
	}

	res, err := h.client.DeleteSale(c.Context(), data)

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

// handlerSaleUpdate function
// @Summary Update Sale
// @Description Update a specific Sale
// @Tags Sale
// @Param id path string true "Sale ID"
// @Param body body domain.UpdateSaleRequest true "Sale Data"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /sale/{id} [put]
func (h *handleSale) handlerSaleUpdate(c *fiber.Ctx) error {
	var body domain.UpdateSaleRequest

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

	data := &pb.UpdateSaleInput{
		Id:      body.ID,
		Name:    body.Name,
		Alamat:  body.Alamat,
		Email:   body.Email,
		Telepon: body.Telepon,
	}

	res, err := h.client.UpdateSale(c.Context(), data)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil mengupdate sale",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

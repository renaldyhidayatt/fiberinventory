package v1

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initSaleGroup(api *fiber.App) {
	routerSale := api.Group("/sale")

	routerSale.Use(middleware.Proctected())

	routerSale.Get("/", h.handlerSaleResults)
	routerSale.Get("/:id", h.handlerSaleResult)
	routerSale.Post("/create", h.handlerSaleCreate)
	routerSale.Post("/update/:id", h.handlerSaleUpdate)
	routerSale.Post("/delete/:id", h.handlerSaleDelete)

}

// handlerSaleCreate function
// @Summary createSale to the application
// @Description Create Sale
// @Tags Sale
// @Accept json
// @Produce json
// @Param body body domain.SaleInput true "Sale information"
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /sale/create [post]
func (h *Handler) handlerSaleCreate(c *fiber.Ctx) error {
	var body domain.SaleInput

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

	res, err := h.services.Sale.Create(&body)

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
// @Summary Get sale results
// @Description Retrieve the results for each sale
// @Tags Sale
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /sale [get]
func (h *Handler) handlerSaleResults(c *fiber.Ctx) error {
	res, err := h.services.Sale.Results()

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
// @Summary Get sale result
// @Description Retrieve the result for a specific sale
// @Tags Sale
// @Param id path string true "Sale ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /sale/{id} [get]
func (h *Handler) handlerSaleResult(c *fiber.Ctx) error {
	var body domain.SaleInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Sale.Result(&body)

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
// @Summary Delete sale
// @Description Delete a specific sale
// @Tags Sale
// @Param id path string true "Sale ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /sale/{id} [delete]
func (h *Handler) handlerSaleDelete(c *fiber.Ctx) error {
	var body domain.SaleInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Sale.Delete(&body)

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
// @Summary Update sale
// @Description Update a specific sale
// @Tags Sale
// @Param id path string true "Sale ID"
// @Param body body domain.SaleInput true "Sale Data"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /sale/{id} [put]
func (h *Handler) handlerSaleUpdate(c *fiber.Ctx) error {
	var body domain.SaleInput
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

	res, err := h.services.Sale.Update(&body)

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

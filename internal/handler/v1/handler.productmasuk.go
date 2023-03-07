package v1

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initProductMasukGroup(api *fiber.App) {
	routerProductMasuk := api.Group("/productmasuk")

	routerProductMasuk.Use(middleware.Proctected())

	routerProductMasuk.Get("/", h.handlerProductMasukResults)
	routerProductMasuk.Get("/:id", h.handlerProductMasukResult)
	routerProductMasuk.Post("/create", h.handlerProductMasukCreate)
	routerProductMasuk.Post("/update/:id", h.handlerProductMasukUpdate)
	routerProductMasuk.Post("/delete/:id", h.handlerProductMasukDelete)

}

// ProductMasukInput function
// @Summary createProductMasuk to the application
// @Description Create ProductMasuk
// @Tags ProductMasuk
// @Accept json
// @Produce json
// @Param productmasuk body domain.ProductMasukInput true "ProductMasuk information"
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productmasuk/create [post]
func (h *Handler) handlerProductMasukCreate(c *fiber.Ctx) error {
	var body domain.ProductMasukInput

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

	res, err := h.services.ProductMasuk.Create(&body)

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
// @Summary Get productmasuk results
// @Description Retrieve the results for each productmasuk
// @Tags ProductMasuk
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productmasuk [get]
func (h *Handler) handlerProductMasukResults(c *fiber.Ctx) error {
	res, err := h.services.ProductMasuk.Results()

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
// @Summary Get productmasuk result
// @Description Retrieve the result for a specific productmasuk
// @Tags ProductMasuk
// @Param id path string true "ProductMasuk ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productmasuk/{id} [get]
func (h *Handler) handlerProductMasukResult(c *fiber.Ctx) error {
	var body domain.ProductMasukInput

	id := c.Params("id")

	body.ID = id

	res, err := h.services.ProductMasuk.Result(&body)

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
// @Summary Delete productMasuk
// @Description Delete a specific productmasuk
// @Tags ProductMasuk
// @Param id path string true "ProductMasuk ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productmasuk/{id} [delete]
func (h *Handler) handlerProductMasukDelete(c *fiber.Ctx) error {
	var body domain.ProductMasukInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.ProductMasuk.Delete(&body)

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
// @Summary Update productmasuk
// @Description Update a specific productmasuk
// @Tags ProductMasuk
// @Param id path string true "ProductMasuk ID"
// @Param body body domain.ProductMasukInput true "ProductMasuk Data"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productmasuk/{id} [put]
func (h *Handler) handlerProductMasukUpdate(c *fiber.Ctx) error {
	var body domain.ProductMasukInput
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
	res, err := h.services.ProductMasuk.Update(&body)

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

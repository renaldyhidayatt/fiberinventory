package v1

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initProductKeluarGroup(api *fiber.App) {
	routerProductKeluar := api.Group("/productkeluar")

	routerProductKeluar.Use(middleware.Proctected())

	routerProductKeluar.Get("/", h.handlerProductKeluarResults)
	routerProductKeluar.Get("/:id", h.handlerProductKeluarResult)
	routerProductKeluar.Post("/create", h.handlerProductKeluarCreate)
	routerProductKeluar.Post("/update/:id", h.handlerProductKeluarUpdate)
	routerProductKeluar.Post("/delete/:id", h.handlerProductKeluarDelete)

}

// handlerProductKeluarCreate function
// @Summary createProductKeluar to the application
// @Description Create ProductKeluar
// @Tags ProductKeluar
// @Accept json
// @Produce json
// @Param productkeluar body domain.ProductKeluarInput true "ProductKeluar information"
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productkeluar/create [post]
func (h *Handler) handlerProductKeluarCreate(c *fiber.Ctx) error {
	var body domain.ProductKeluarInput

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

	res, err := h.services.ProductKeluar.Create(&body)

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
// @Summary Get productkeluar results
// @Description Retrieve the results for each productkeluar
// @Tags ProductKeluar
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productkeluar [get]
func (h *Handler) handlerProductKeluarResults(c *fiber.Ctx) error {
	res, err := h.services.ProductKeluar.Results()

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
// @Summary Get productkeluar result
// @Description Retrieve the result for a specific productkeluar
// @Tags ProductKeluar
// @Param id path string true "ProductKeluar ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productkeluar/{id} [get]
func (h *Handler) handlerProductKeluarResult(c *fiber.Ctx) error {
	var body domain.ProductKeluarInput

	id := c.Params("id")

	body.ID = id

	res, err := h.services.ProductKeluar.Result(&body)

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
// @Summary Delete productkeluar
// @Description Delete a specific productkeluar
// @Tags ProductKeluar
// @Param id path string true "ProductKeluar ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productkeluar/{id} [delete]
func (h *Handler) handlerProductKeluarDelete(c *fiber.Ctx) error {
	var body domain.ProductKeluarInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.ProductKeluar.Delete(&body)

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
// @Summary Update productkeluar
// @Description Update a specific productkeluar
// @Tags ProductKeluar
// @Param id path string true "ProductKeluar ID"
// @Param body body domain.ProductKeluarInput true "ProductKeluar Data"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /productkeluar/{id} [put]
func (h *Handler) handlerProductKeluarUpdate(c *fiber.Ctx) error {
	var body domain.ProductKeluarInput
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

	res, err := h.services.ProductKeluar.Update(&body)

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

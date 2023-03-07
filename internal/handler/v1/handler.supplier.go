package v1

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initSupplierGroup(api *fiber.App) {
	routerSupplier := api.Group("/supplier")

	routerSupplier.Use(middleware.Proctected())

	routerSupplier.Get("/", h.handlerSupplierResults)
	routerSupplier.Get("/:id", h.handlerSupplierResult)
	routerSupplier.Post("/create", h.handlerSupplierCreate)
	routerSupplier.Post("/update/:id", h.handlerSupplierUpdate)
	routerSupplier.Post("/delete/:id", h.handlerSupplierDelete)

}

// handlerSupplierCreate function
// @Summary createSupplier to the application
// @Description Create Supplier
// @Tags Supplier
// @Accept json
// @Produce json
// @Param body body domain.SupplierInput true "Supplier information"
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /supplier/create [post]
func (h *Handler) handlerSupplierCreate(c *fiber.Ctx) error {
	var body domain.SupplierInput

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

	res, err := h.services.Supplier.Create(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil membuat supplier",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// handlerSupplierResults function
// @Summary Get supplier results
// @Description Retrieve the results for each supplier
// @Tags Supplier
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /supplier [get]
func (h *Handler) handlerSupplierResults(c *fiber.Ctx) error {
	res, err := h.services.Supplier.Results()

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

// handlerSupplierResult function
// @Summary Get supplier result
// @Description Retrieve the result for a specific supplier
// @Tags Supplier
// @Param id path string true "Supplier ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /supplier/{id} [get]
func (h *Handler) handlerSupplierResult(c *fiber.Ctx) error {
	var body domain.SupplierInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Supplier.Result(&body)

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

// handlerSupplierDelete function
// @Summary Delete supplier
// @Description Delete a specific supplier
// @Tags Supplier
// @Param id path string true "Supplier ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /supplier/{id} [delete]
func (h *Handler) handlerSupplierDelete(c *fiber.Ctx) error {
	var body domain.SupplierInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Supplier.Delete(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil menghapus supplier",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

// handlerSupplierUpdate function
// @Summary Update supplier
// @Description Update a specific supplier
// @Tags Supplier
// @Param id path string true "Supplier ID"
// @Param body body domain.SupplierInput true "Category Data"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /sale/{id} [put]
func (h *Handler) handlerSupplierUpdate(c *fiber.Ctx) error {
	var body domain.SupplierInput

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

	res, err := h.services.Supplier.Update(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(domain.ErrorMessage{
			Error:      true,
			Message:    err.Error(),
			StatusCode: fiber.StatusBadRequest,
		})
	}

	return c.Status(fiber.StatusOK).JSON(domain.Response{
		Message:    "berhasil mengupdate supplier",
		Data:       res,
		StatusCode: fiber.StatusOK,
	})
}

package api

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"
	"fiberinventory/internal/pb"

	"github.com/gofiber/fiber/v2"
)

type handleSupplier struct {
	client pb.SupplierServiceClient
}

func NewHandleSupplier(client pb.SupplierServiceClient, router *fiber.App) {
	h := handleSupplier{
		client: client,
	}

	routerSupplier := router.Group("/api/supplier")

	routerSupplier.Use(middleware.Proctected())

	routerSupplier.Get("/hello", h.handlerSupplierHello)

	routerSupplier.Get("/", h.handlerSupplierResults)
	routerSupplier.Get("/:id", h.handlerSupplierResult)
	routerSupplier.Post("/create", h.handlerSupplierCreate)
	routerSupplier.Post("/update/:id", h.handlerSupplierUpdate)
	routerSupplier.Post("/delete/:id", h.handlerSupplierDelete)
}

// handlerSupplierHello function
// @Summary Greet the Category
// @Description Return a greeting message to the supplier
// @Tags Supplier
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /api/supplier/hello [get]
func (h *handleSupplier) handlerSupplierHello(c *fiber.Ctx) error {
	return c.SendString("Handler Supplier")
}

// handlerSupplierCreate function
// @Summary createSupplier to the application
// @Description Create Supplier
// @Tags Supplier
// @Accept json
// @Produce json
// @Param body body domain.CreateSupplierRequest true "Supplier information"
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /supplier/create [post]
func (h *handleSupplier) handlerSupplierCreate(c *fiber.Ctx) error {
	var body domain.CreateSupplierRequest

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

	data := &pb.CreateSupplierInput{
		Name:    body.Name,
		Alamat:  body.Alamat,
		Email:   body.Email,
		Telepon: body.Telepon,
	}

	res, err := h.client.CreateSupplier(c.Context(), data)

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
func (h *handleSupplier) handlerSupplierResults(c *fiber.Ctx) error {
	res, err := h.client.GetSuppliers(c.Context(), &pb.SuppliersRequest{})

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
// @Router /api/supplier/{id} [get]
func (h *handleSupplier) handlerSupplierResult(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.SupplierRequest{
		Id: id,
	}

	res, err := h.client.GetSupplier(c.Context(), data)

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
// @Router /api/supplier/{id} [delete]
func (h *handleSupplier) handlerSupplierDelete(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.SupplierRequest{
		Id: id,
	}

	res, err := h.client.DeleteSupplier(c.Context(), data)

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
// @Param body body domain.UpdateSupplierRequest true "Supplier Data"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /api/supplier/{id} [put]
func (h *handleSupplier) handlerSupplierUpdate(c *fiber.Ctx) error {
	var body domain.UpdateSupplierRequest

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

	data := &pb.UpdateSupplierInput{
		Id:      body.ID,
		Name:    body.Name,
		Alamat:  body.Alamat,
		Email:   body.Email,
		Telepon: body.Telepon,
	}

	res, err := h.client.UpdateSupplier(c.Context(), data)

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

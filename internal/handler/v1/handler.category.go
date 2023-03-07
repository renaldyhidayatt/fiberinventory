package v1

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"

	"github.com/gofiber/fiber/v2"
)

func (h *Handler) initCategoryGroup(api *fiber.App) {
	routerCategory := api.Group("/category")

	routerCategory.Use(middleware.Proctected())

	routerCategory.Get("/hello", h.handlerCategoryHello)
	routerCategory.Post("/create", h.handlerCategoryCreate)
	routerCategory.Put("/update/:id", h.handlerCategoryUpdate)
	routerCategory.Delete("/delete/:id", h.handlerCategoryDelete)
	routerCategory.Get("/", h.handleCategoryResults)
	routerCategory.Get("/:id", h.handlerCategoryResult)
}

// handlerCategoryHello function
// @Summary Greet the Category
// @Description Return a greeting message to the category
// @Tags Category
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /category/hello [get]
func (h *Handler) handlerCategoryHello(c *fiber.Ctx) error {
	return c.SendString("Handler Category")
}

// handlerCategoryCreate function
// @Summary createCategory to the application
// @Description Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param body body domain.CategoryInput true "Category information"
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /category/create [post]
func (h *Handler) handlerCategoryCreate(c *fiber.Ctx) error {
	var body domain.CategoryInput
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

	res, err := h.services.Category.Create(&body)

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

// handleCategoryResults function
// @Summary Get category results
// @Description Retrieve the results for each category
// @Tags Category
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /category [get]
func (h *Handler) handleCategoryResults(c *fiber.Ctx) error {
	res, err := h.services.Category.Results()

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

// handlerCategoryResult function
// @Summary Get category result
// @Description Retrieve the result for a specific category
// @Tags Category
// @Param id path string true "Category ID"
// @Produce json
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /category/{id} [get]
func (h *Handler) handlerCategoryResult(c *fiber.Ctx) error {
	var body domain.CategoryInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Category.Result(&body)

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

// handlerCategoryDelete function
// @Summary Delete category
// @Description Delete a specific category
// @Tags Category
// @Param id path string true "Category ID"
// @Security BearerAuth
// @Produce json
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /category/{id} [delete]
func (h *Handler) handlerCategoryDelete(c *fiber.Ctx) error {
	var body domain.CategoryInput

	id := c.Params("id")
	body.ID = id

	res, err := h.services.Category.Delete(&body)

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

// handlerCategoryUpdate function
// @Summary Update category
// @Description Update a specific category
// @Tags Category
// @Param id path string true "Category ID"
// @Param body body domain.CategoryInput true "Category Data"
// @Security BearerAuth
// @Produce json
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /category/{id} [put]
func (h *Handler) handlerCategoryUpdate(c *fiber.Ctx) error {
	var body domain.CategoryInput
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

	res, err := h.services.Category.Update(&body)

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

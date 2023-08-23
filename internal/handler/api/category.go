package api

import (
	"fiberinventory/internal/domain"
	"fiberinventory/internal/middleware"
	"fiberinventory/internal/pb"

	"github.com/gofiber/fiber/v2"
)

type handleCategory struct {
	client pb.CategoryServiceClient
}

func NewHandlerCategory(client pb.CategoryServiceClient, router *fiber.App) {
	handlerCategory := handleCategory{
		client: client,
	}

	routerCategory := router.Group("/api/category")
	routerCategory.Use(middleware.Proctected())

	routerCategory.Get("/hello", handlerCategory.handlerCategoryHello)
	routerCategory.Post("/create", handlerCategory.handlerCategoryCreate)
	routerCategory.Put("/update/:id", handlerCategory.handlerCategoryUpdate)
	routerCategory.Delete("/delete/:id", handlerCategory.handlerCategoryDelete)
	routerCategory.Get("/", handlerCategory.handleCategoryResults)
	routerCategory.Get("/:id", handlerCategory.handlerCategoryResult)

}

// handlerCategoryHello function
// @Summary Greet the Category
// @Description Return a greeting message to the category
// @Tags Category
// @Produce plain
// @Success 200 {string} string "OK"
// @Router /category/hello [get]
func (h *handleCategory) handlerCategoryHello(c *fiber.Ctx) error {
	return c.SendString("Handler Category")
}

// handlerCategoryCreate function
// @Summary createCategory to the application
// @Description Create Category
// @Tags Category
// @Accept json
// @Produce json
// @Param body body domain.CreateCategoryRequest true "Category information"
// @Security BearerAuth
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /category/create [post]
func (h *handleCategory) handlerCategoryCreate(c *fiber.Ctx) error {
	var body domain.CreateCategoryRequest

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

	data := &pb.CreateCategoryInput{
		Name: body.Name,
	}

	res, err := h.client.CreateCategory(c.Context(), data)

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
func (h *handleCategory) handleCategoryResults(c *fiber.Ctx) error {

	res, err := h.client.GetCategories(c.Context(), &pb.CategoriesRequest{})

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
func (h *handleCategory) handlerCategoryResult(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.CategoryRequest{
		Id: id,
	}

	res, err := h.client.GetCategory(c.Context(), data)

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

// handlerCategoryUpdate function
// @Summary Update category
// @Description Update a specific category
// @Tags Category
// @Param id path string true "Category ID"
// @Param body body domain.UpdateCategoryRequest true "Category Data"
// @Security BearerAuth
// @Produce json
// @Success 200 {object} domain.Response
// @Failure 400 {object} domain.ErrorMessage
// @Router /category/{id} [put]
func (h *handleCategory) handlerCategoryUpdate(c *fiber.Ctx) error {
	var body domain.UpdateCategoryRequest

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

	data := &pb.UpdateCategoryInput{
		Id:   body.ID,
		Name: body.Name,
	}

	res, err := h.client.UpdateCategory(c.Context(), data)

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
func (h *handleCategory) handlerCategoryDelete(c *fiber.Ctx) error {

	id := c.Params("id")

	data := &pb.CategoryRequest{
		Id: id,
	}

	res, err := h.client.DeleteCategory(c.Context(), data)

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

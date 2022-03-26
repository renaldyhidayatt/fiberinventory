package routes

import (
	"fiberinventory/handler"
	"fiberinventory/middleware"
	"fiberinventory/repository"
	"fiberinventory/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouteMerchant(db *gorm.DB, router *fiber.App) {
	repository := repository.NewRepositoryCategory(db)
	service := service.NewServiceCategory(repository)
	handler := handler.NewHandlerCategory(service)

	route := router.Group("/api/category")

	route.Use(middleware.Proctected())
	route.Get("/hello", handler.HandlerHello)
	route.Post("/create", handler.HandlerCreate)
	route.Put("/update/:id", handler.HandlerUpdate)
	route.Delete("/delete/:id", handler.HandlerDelete)
	route.Get("/", handler.HandlerResults)
	route.Get("/:id", handler.HandlerResult)
}

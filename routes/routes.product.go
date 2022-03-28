package routes

import (
	"fiberinventory/handler"
	"fiberinventory/middleware"
	"fiberinventory/repository"
	"fiberinventory/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouteProduct(db *gorm.DB, router *fiber.App) {
	repository := repository.NewRepositoryProduct(db)
	service := service.NewServiceProduct(repository)
	handler := handler.NewHandlerProduct(service)

	route := router.Group("/api/product")

	route.Use(middleware.Proctected())

	route.Get("/hello", handler.HandlerHello)
	router.Get("/", handler.HandlerResults)
	router.Get("/:id", handler.HandlerResult)
	router.Post("/create", handler.HandlerCreate)
	router.Post("/update/:id", handler.HandlerUpdate)
	router.Post("/delete/:id", handler.HandlerDelete)
}

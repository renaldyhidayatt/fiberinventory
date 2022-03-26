package routes

import (
	"fiberinventory/handler"
	"fiberinventory/middleware"
	"fiberinventory/repository"
	"fiberinventory/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouteSupplier(db *gorm.DB, router *fiber.App) {
	repository := repository.NewRepositorySupplier(db)
	service := service.NewServiceSupplier(repository)
	handler := handler.NewHandlerSupplier(service)

	route := router.Group("/api/supplier")

	route.Use(middleware.Proctected())

	route.Get("/hello", handler.HandlerHello)
	router.Get("/", handler.HandlerResults)
	router.Get("/:id", handler.HandlerResults)
	router.Post("/create", handler.HandlerCreate)
	router.Post("/update/:id", handler.HandlerUpdate)
	router.Post("/delete/:id", handler.HandlerDelete)
}

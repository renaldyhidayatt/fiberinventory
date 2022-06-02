package routes

import (
	"fiberinventory/handler"
	"fiberinventory/middleware"
	"fiberinventory/repository"
	"fiberinventory/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouteProductKeluar(db *gorm.DB, router *fiber.App) {
	repository := repository.NewRepositoryProductKeluar(db)
	service := service.NewServiceProductKeluar(repository)
	handler := handler.NewHandlerProductKeluar(service)

	route := router.Group("/api/productkeluar")

	route.Use(middleware.Proctected())
	router.Get("/", handler.HandlerResults)
	router.Get("/:id", handler.HandlerResult)
	router.Post("/create", handler.HandlerCreate)
	router.Post("/update/:id", handler.HandlerUpdate)
	router.Post("/delete/:id", handler.HandlerDelete)

}

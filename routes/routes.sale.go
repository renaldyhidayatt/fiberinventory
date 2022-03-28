package routes

import (
	"fiberinventory/handler"
	"fiberinventory/middleware"
	"fiberinventory/repository"
	"fiberinventory/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouteSale(db *gorm.DB, router *fiber.App) {
	repository := repository.NewRepositorySale(db)
	service := service.NewServiceSale(repository)
	handler := handler.NewHandlerSale(service)

	route := router.Group("/api/sale")

	route.Use(middleware.Proctected())

	route.Get("/hello", handler.HandlerHello)
	router.Get("/", handler.HandlerResults)
	router.Get("/:id", handler.HandlerResult)
	router.Post("/create", handler.HandlerCreate)
	router.Post("/update/:id", handler.HandlerUpdate)
	router.Post("/delete/:id", handler.HandlerDelete)
}

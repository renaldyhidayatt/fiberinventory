package routes

import (
	"fiberinventory/handler"
	"fiberinventory/middleware"
	"fiberinventory/repository"
	"fiberinventory/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouteCustomer(db *gorm.DB, router *fiber.App) {
	repository := repository.NewRepositoryCustomer(db)
	service := service.NewServiceCustomer(repository)
	handler := handler.NewHandlerCustomer(service)

	route := router.Group("/api/customer")

	route.Use(middleware.Proctected())

	route.Get("/hello", handler.HandlerHello)
	router.Get("/", handler.HandlerResults)
	router.Get("/:id", handler.HandlerResults)
	router.Post("/create", handler.HandlerCreate)
	router.Post("/update/:id", handler.HandlerUpdate)
	router.Post("/delete/:id", handler.HandlerDelete)
}

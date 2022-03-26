package routes

import (
	"fiberinventory/handler"
	"fiberinventory/repository"
	"fiberinventory/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRouteAuth(db *gorm.DB, router *fiber.App) {
	repository := repository.NewRepositoryUser(db)
	service := service.NewServiceAuth(repository)
	handler := handler.NewHandlerAuth(service)

	route := router.Group("/api/auth")

	route.Get("/hello", handler.HandlerHello)
	route.Post("/register", handler.HandlerRegister)
	route.Post("/login", handler.HandlerLogin)
}

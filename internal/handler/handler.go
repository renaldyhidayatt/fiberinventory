package handler

import (
	v1 "fiberinventory/internal/handler/v1"
	"fiberinventory/internal/service"

	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{
		services: services,
	}
}

func (h *Handler) Init() *fiber.App {

	router := fiber.New()

	router.Use(logger.New())

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	h.InitApi(router)

	return router
}

func (h *Handler) InitApi(router *fiber.App) {
	handlerV1 := v1.NewHandler(h.services)

	handlerV1.Init(router)
}

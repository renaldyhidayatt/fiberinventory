package handler

import (
	v1 "fiberinventory/internal/handler/v1"
	"fiberinventory/internal/service"
	"time"

	_ "fiberinventory/docs"

	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"

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

const idleTimeout = 5 * time.Second

func (h *Handler) Init() *fiber.App {

	router := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	router.Use(logger.New())

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})
	router.Get("/docs/*", swagger.HandlerDefault)

	h.InitApi(router)

	return router
}

func (h *Handler) InitApi(router *fiber.App) {
	handlerV1 := v1.NewHandler(h.services)

	handlerV1.Init(router)
}

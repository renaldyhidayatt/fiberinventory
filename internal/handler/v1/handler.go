package v1

import (
	"fiberinventory/internal/service"

	"github.com/gofiber/fiber/v2"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) Init(api *fiber.App) {

	h.initAuthGroup(api)
	h.initCategoryGroup(api)
	h.initCustomer(api)
	h.initProductGroup(api)
	h.initProductKeluarGroup(api)
	h.initProductMasukGroup(api)
	h.initSaleGroup(api)
	h.initSupplierGroup(api)

}

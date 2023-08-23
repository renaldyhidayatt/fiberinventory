package api

import (
	_ "fiberinventory/docs"
	"fiberinventory/internal/pb"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/swagger"
)

type handler struct {
	auth          pb.AuthServiceClient
	category      pb.CategoryServiceClient
	customer      pb.CustomerServiceClient
	product       pb.ProductServiceClient
	productKeluar pb.ProductKeluarServiceClient
	productMasuk  pb.ProductMasukServiceClient
	sale          pb.SaleServiceClient
	supplier      pb.SupplierServiceClient
	user          pb.UserServiceClient
}

func NewHandler(auth pb.AuthServiceClient, category pb.CategoryServiceClient, customer pb.CustomerServiceClient, product pb.ProductServiceClient, productKeluar pb.ProductKeluarServiceClient, productMasuk pb.ProductMasukServiceClient, sale pb.SaleServiceClient, supplier pb.SupplierServiceClient, user pb.UserServiceClient) *handler {
	return &handler{
		auth:          auth,
		category:      category,
		customer:      customer,
		product:       product,
		productKeluar: productKeluar,
		productMasuk:  productMasuk,
		sale:          sale,
		supplier:      supplier,
		user:          user,
	}
}

const idleTimeout = 5 * time.Second

func (h *handler) Init() *fiber.App {
	router := fiber.New(fiber.Config{
		IdleTimeout: idleTimeout,
	})

	router.Use(logger.New())

	router.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello World")
	})

	router.Get("/docs/*", swagger.HandlerDefault)

	NewHandlerAuth(h.auth, router)
	NewHandlerCategory(h.category, router)
	NewHandlerCustomer(h.customer, router)
	NewHandleProduct(h.product, router)
	NewHandleProductKeluar(h.productKeluar, router)
	NewHandleProductMasuk(h.productMasuk, router)
	NewHandlerSale(h.sale, router)
	NewHandleSupplier(h.supplier, router)
	NewHandlerUser(h.user, router)

	return router
}

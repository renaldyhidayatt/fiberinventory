package routes

import (
	"fiberinventory/handler"
	"fiberinventory/repository"
	"fiberinventory/service"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func NewRoute(db *gorm.DB, router *fiber.App) {
	// Authentication

	repositoryAuth := repository.NewRepositoryUser(db)
	serviceAuth := service.NewServiceAuth(repositoryAuth)
	handlerAuth := handler.NewHandlerAuth(serviceAuth)

	// Category
	repositoryCategory := repository.NewRepositoryCategory(db)
	serviceCategory := service.NewServiceCategory(repositoryCategory)
	handlerCategory := handler.NewHandlerCategory(serviceCategory)

	// Product
	repositoryProduct := repository.NewRepositoryProduct(db)
	serviceProduct := service.NewServiceProduct(repositoryProduct)
	handlerProduct := handler.NewHandlerProduct(serviceProduct)

	// ProductKeluar
	repositoryProductKeluar := repository.NewRepositoryProductKeluar(db)
	serviceProductKeluar := service.NewServiceProductKeluar(repositoryProductKeluar)
	handlerProductKeluar := handler.NewHandlerProductKeluar(serviceProductKeluar)

	// Product Masuk
	repositoryProductMasuk := repository.NewRepositoryProductMasuk(db)
	serviceProductMasuk := service.NewServiceProductMasuk(repositoryProductMasuk)
	handlerProductMasuk := handler.NewHandlerProductMasuk(serviceProductMasuk)

	// Sale
	repositorySale := repository.NewRepositorySale(db)
	serviceSale := service.NewServiceSale(repositorySale)
	handlerSale := handler.NewHandlerSale(serviceSale)

	// Supplier
	repositorySupplier := repository.NewRepositorySupplier(db)
	serviceSupplier := service.NewServiceSupplier(repositorySupplier)
	handlerSupplier := handler.NewHandlerSupplier(serviceSupplier)

	routerAuth := router.Group("/api/auth")

	routerCustomer := router.Group("/api/group")

	routerCategory := router.Group("/api/category")

	routerProduct := router.Group("/api/product")

	routerProductKeluar := router.Group("/api/productkeluar")

	routerProductMasuk := router.Group("/api/productmasuk")

	routerSale := router.Group("/api/sale")

	routeSupplier := router.Group("/api/supplier")

	routerAuth.Get("/hello", handlerAuth.HandlerHello)
	routerAuth.Post("/register", handlerAuth.HandlerRegister)
	routerAuth.Post("/login", handlerAuth.HandlerLogin)

	routerCategory.Get("/hello", handlerCategory.HandlerHello)
	routerCategory.Post("/create", handlerCategory.HandlerCreate)
	routerCategory.Put("/update/:id", handlerCategory.HandlerUpdate)
	routerCategory.Delete("/delete/:id", handlerCategory.HandlerDelete)
	routerCategory.Get("/", handlerCategory.HandlerResults)
	routerCategory.Get("/:id", handlerCategory.HandlerResult)

	routerCustomer.Get("/hello", handlerCategory.HandlerHello)
	routerCustomer.Get("/", handlerCategory.HandlerResults)
	routerCustomer.Get("/:id", handlerCategory.HandlerResults)
	routerCustomer.Post("/create", handlerCategory.HandlerCreate)
	routerCustomer.Post("/update/:id", handlerCategory.HandlerUpdate)
	routerCustomer.Post("/delete/:id", handlerCategory.HandlerDelete)

	routerProduct.Get("/hello", handlerProduct.HandlerHello)
	routerProduct.Get("/", handlerProduct.HandlerResults)
	routerProduct.Get("/:id", handlerProduct.HandlerResult)
	routerProduct.Post("/create", handlerProduct.HandlerCreate)
	routerProduct.Post("/update/:id", handlerProduct.HandlerUpdate)
	routerProduct.Post("/delete/:id", handlerProduct.HandlerDelete)

	routerProductKeluar.Get("/", handlerProductKeluar.HandlerResults)
	routerProductKeluar.Get("/:id", handlerProductKeluar.HandlerResult)
	routerProductKeluar.Post("/create", handlerProductKeluar.HandlerCreate)
	routerProductKeluar.Post("/update/:id", handlerProductKeluar.HandlerUpdate)
	routerProductKeluar.Post("/delete/:id", handlerProductKeluar.HandlerDelete)

	routerProductMasuk.Get("/hello", handlerProductMasuk.HandlerHello)
	routerProductMasuk.Get("/", handlerProductMasuk.HandlerResults)
	routerProductMasuk.Get("/:id", handlerProductMasuk.HandlerResult)
	routerProductMasuk.Post("/create", handlerProductMasuk.HandlerCreate)
	routerProductMasuk.Post("/update/:id", handlerProductMasuk.HandlerUpdate)
	routerProductMasuk.Post("/delete/:id", handlerProductMasuk.HandlerDelete)

	routerSale.Get("/hello", handlerSale.HandlerHello)
	routerSale.Get("/", handlerSale.HandlerResults)
	routerSale.Get("/:id", handlerSale.HandlerResult)
	routerSale.Post("/create", handlerSale.HandlerCreate)
	routerSale.Post("/update/:id", handlerSale.HandlerUpdate)
	routerSale.Post("/delete/:id", handlerSale.HandlerDelete)

	routeSupplier.Get("/hello", handlerSupplier.HandlerHello)
	routeSupplier.Get("/", handlerSupplier.HandlerResults)
	routeSupplier.Get("/:id", handlerSupplier.HandlerResults)
	routeSupplier.Post("/create", handlerSupplier.HandlerCreate)
	routeSupplier.Post("/update/:id", handlerSupplier.HandlerUpdate)
	routeSupplier.Post("/delete/:id", handlerSupplier.HandlerDelete)
}

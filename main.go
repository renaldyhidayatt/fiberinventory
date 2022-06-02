package main

import (
	"fiberinventory/config"
	"fiberinventory/routes"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func main() {
	app := fiber.New()
	db := config.SetupDatabase()

	app.Get("/", helloWorld)

	routes.NewRouteAuth(db, app)
	routes.NewRouteMerchant(db, app)
	routes.NewRouteCustomer(db, app)
	routes.NewRouteSupplier(db, app)
	routes.NewRouteSale(db, app)
	routes.NewRouteProduct(db, app)
	routes.NewRouteProductMasuk(db, app)
	routes.NewRouteProductKeluar(db, app)

	app.Listen(":5000")
}

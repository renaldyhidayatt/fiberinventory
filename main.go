package main

import (
	"fiberinventory/config"
	"fiberinventory/middleware"
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

	app.Use(middleware.Proctected())

	routes.NewRoute(db, app)

	app.Listen(":5000")
}

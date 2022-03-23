package main

import (
	"fiberinventory/config"

	"github.com/gofiber/fiber/v2"
)

func helloWorld(c *fiber.Ctx) error {
	return c.SendString("Hello World")
}

func main() {
	app := fiber.New()
	config.SetupDatabase()

	app.Get("/", helloWorld)

	app.Listen(":5000")
}

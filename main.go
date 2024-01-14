package main

import (
	"blog/model"
	"blog/router"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	app.Use(recover.New())

	router.Router(app)
	app.Use("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
		})
	})

	app.Static("/")
	model.Init()
	log.Fatal(app.Listen(":3000"))
}

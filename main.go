package main

import (
	"blog/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	router.Router(app)
	app.All("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
		})
	})

	log.Fatal(app.Listen(":3000"))
}

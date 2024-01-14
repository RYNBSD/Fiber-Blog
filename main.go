package main

import (
	"blog/model"
	"blog/router"
	"blog/util"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New()
	// If error is Internal use panic
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
		StackTraceHandler: func(c *fiber.Ctx, err any) {
			status := 0
			message := ""

			switch e := err.(type) {
			case error:
				status = fiber.StatusInternalServerError
				message = e.Error()
			case fiber.Error:
				status = e.Code
				message = e.Message
			case string:
				status = fiber.StatusInternalServerError
				message = e
			default:
				status = fiber.StatusInternalServerError
				message = fmt.Sprintf("%v", err)
			}

			
		},
	}))

	router.Router(app)
	app.Use("*", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"success": false,
		})
	})

	app.Static("/", util.PublicDir())
	model.Init()
	log.Fatal(app.Listen(":3000"))
}

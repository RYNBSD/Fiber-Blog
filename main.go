package main

import (
	"blog/model"
	"blog/router"
	"blog/util"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func main() {
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// status := 0
			// message := ""

			// switch e := err.(type) {
			// case error:
			// 	status = fiber.StatusInternalServerError
			// 	message = e.Error()
			// case fiber.Error:
			// 	status = e.Code
			// 	message = e.Message
			// case string:
			// 	status = fiber.StatusInternalServerError
			// 	message = e
			// default:
			// 	status = fiber.StatusInternalServerError
			// 	message = fmt.Sprintf("%v", e)
			// }
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"success": false,
			})
		},
	})
	app.Use(cors.New())
	app.Use(limiter.New())
	app.Use(logger.New())
	app.Use(helmet.New())
	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	router.Router(app)
	app.Use("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	app.Static("/", util.PublicDir())
	model.Init()
	log.Fatal(app.Listen(":3000"))
}

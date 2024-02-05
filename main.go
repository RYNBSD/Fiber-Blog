package main

import (
	"blog/model"
	"blog/router"
	"blog/util"
	"errors"
	"fmt"
	"log"
	"net/http"

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
			code := fiber.StatusInternalServerError
			message := ""
			
			var fiberError *fiber.Error
			if errors.As(err, &fiberError) {
				code = fiberError.Code
			}

			if code >= 500 {
				e := model.Error{
					Status: int16(code),
					Message: err.Error(),
				}

				e.CreateError()
				message = http.StatusText(code)
			} else {
				errorMessage := err.Error()
				if len(errorMessage) > 0 {
					message = errorMessage
				} else {
					message = http.StatusText(code)
				}
			}
			
			return c.Status(code).JSON(fiber.Map{
				"success": false,
				"message": message,
			})
		},
	})
	app.Use(cors.New())
	app.Use(limiter.New())
	app.Use(logger.New())
	app.Use(helmet.New())
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
				message = fmt.Sprintf("%v", e)
			}

			e := model.Error{
				Status: int16(status),
				Message: message,
			}
			e.CreateError()
		},
	}))
	router.Router(app)
	app.Static("/", util.PublicDir())
	app.Use("*", func(c *fiber.Ctx) error {
		return c.SendStatus(fiber.StatusNotFound)
	})

	model.Init()
	log.Fatal(app.Listen(":3000"))
}

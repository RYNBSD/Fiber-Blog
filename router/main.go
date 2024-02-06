package router

import "github.com/gofiber/fiber/v2"

func Router(app *fiber.App) {
	API(app)
	Security(app)
}

package router

import (
	"blog/middleware"

	"github.com/gofiber/fiber/v2"
)

var security fiber.Router

func Security(app *fiber.App) {
	security = app.Group("/security")
	access()
}

func access() {
	access := security.Group("/access")

	access.Post("/email", middleware.HasUserUnregistered)
}

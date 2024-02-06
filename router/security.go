package router

import (
	"blog/controller/security"
	"blog/middleware"

	"github.com/gofiber/fiber/v2"
)

var security_router fiber.Router

func Security(app *fiber.App) {
	security_router = app.Group("/security")
	access()
}

func access() {
	access := security_router.Group("/access")

	access.Post("/email", middleware.HasUserUnregistered, security.AccessEmail)
}

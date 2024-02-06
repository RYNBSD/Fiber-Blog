package router

import (
	"blog/controller"
	"blog/middleware"

	"github.com/gofiber/fiber/v2"
)

var api fiber.Router

func Router(app *fiber.App) {
	api = app.Group("/api")
	auth()
	user()
	blog()
}

func auth() {
	auth := api.Group("/auth")

	auth.Post("/sign-up", middleware.HasUserUnregistered, controller.SignUp)
	auth.Post("/sign-in", middleware.HasUserUnregistered, controller.SignIn)
	auth.Post("/sign-out", middleware.HasUserUnregistered, controller.SignOut)
	auth.Post("/me", middleware.HasUserUnregistered, controller.Me)
}

func user() {

}

func blog() {

}

package router

import (
	"blog/controller"

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

	auth.Post("/sign-up", controller.SignUp)
	auth.Post("/sign-in", controller.SignIn)
	auth.Post("/sign-out", controller.SignOut)
	auth.Post("/me", controller.Me)
}

func user() {

}

func blog() {

}
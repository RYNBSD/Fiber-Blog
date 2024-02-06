package router

import (
	"blog/controller"
	"blog/middleware"

	"github.com/gofiber/fiber/v2"
)

var api fiber.Router

func API(app *fiber.App) {
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

	auth.Put("/forgot-password", middleware.HasUserUnregistered)
}

func user() {
	user := api.Group("/user")
	
	user.Get("/:userId/info", middleware.HasUserUnregistered)
	user.Get("/:userId/blogs", middleware.HasUserUnregistered)
	
	user.Put("/", middleware.HasUserRegistered)
	user.Delete("/", middleware.HasUserRegistered)
}

func blog() {
	blog := api.Group("/blog")

	blog.Get("/all", middleware.HasUserUnregistered)
	blog.Get("/:blogId", middleware.HasUserUnregistered)
	blog.Get("/:blogId/likes", middleware.HasUserUnregistered)
	blog.Get("/:blogId/comments", middleware.HasUserUnregistered)

	blog.Post("/", middleware.HasUserRegistered)
	blog.Post("/:blogId/comment", middleware.HasUserRegistered)
	
	blog.Put("/:blogId", middleware.HasUserRegistered)
	blog.Put("/:blogId/commentId", middleware.HasUserRegistered)
	
	blog.Delete("/:blogId", middleware.HasUserRegistered)
	blog.Delete("/:blogId/:commentId", middleware.HasUserRegistered)
}

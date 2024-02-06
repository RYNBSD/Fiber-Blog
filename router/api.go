package router

import (
	"blog/controller/api"
	"blog/middleware"

	"github.com/gofiber/fiber/v2"
)

var api_router fiber.Router

func API(app *fiber.App) {
	api_router = app.Group("/api")
	auth()
	user()
	blog()
}

func auth() {
	auth := api_router.Group("/auth")

	auth.Post("/sign-up", middleware.HasUserUnregistered, api.SignUp)
	auth.Post("/sign-in", middleware.HasUserUnregistered, api.SignIn)
	auth.Post("/sign-out", middleware.HasUserUnregistered, api.SignOut)
	auth.Post("/me", middleware.HasUserUnregistered, api.Me)

	auth.Put("/forgot-password", middleware.HasUserUnregistered)
}

func user() {
	user := api_router.Group("/user")
	
	user.Get("/:userId/info", middleware.HasUserUnregistered)
	user.Get("/:userId/blogs", middleware.HasUserUnregistered)
	
	user.Put("/", middleware.HasUserRegistered)
	user.Delete("/", middleware.HasUserRegistered)
}

func blog() {
	blog := api_router.Group("/blog")

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

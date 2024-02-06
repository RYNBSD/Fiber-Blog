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

	auth.Put("/forgot-password", middleware.HasUserUnregistered, api.ForgotPassword)
}

func user() {
	user := api_router.Group("/user")

	user.Get("/:userId/info", middleware.HasUserUnregistered, api.Info)
	user.Get("/:userId/blogs", middleware.HasUserUnregistered, api.Blogs)

	user.Put("/", middleware.HasUserRegistered, api.Update)
	user.Delete("/", middleware.HasUserRegistered, api.Delete)
}

func blog() {
	blog := api_router.Group("/blog")

	blog.Get("/all", middleware.HasUserUnregistered, api.All)
	blog.Get("/:blogId", middleware.HasUserUnregistered, api.Blog)
	blog.Get("/:blogId/likes", middleware.HasUserUnregistered, api.Likes)
	blog.Get("/:blogId/comments", middleware.HasUserUnregistered, api.Comments)

	blog.Post("/", middleware.HasUserRegistered, api.CreateBlog)
	blog.Post("/:blogId/comment", middleware.HasUserRegistered, api.CreateComment)

	blog.Put("/:blogId", middleware.HasUserRegistered, api.UpdateBlog)
	blog.Put("/:blogId/commentId", middleware.HasUserRegistered, api.UpdateComment)

	blog.Delete("/:blogId", middleware.HasUserRegistered, api.DeleteBlog)
	blog.Delete("/:blogId/:commentId", middleware.HasUserRegistered, api.DeleteComment)
}

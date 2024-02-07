package api

import (
	"blog/model"

	"github.com/gofiber/fiber/v2"
)

func All(c *fiber.Ctx) error {
	blog := model.Blog{}
	blogs := blog.SelectBlogs()

	status := fiber.StatusOK
	if len(blogs) == 0 {
		status = fiber.StatusNoContent
	}

	return c.Status(status).JSON(fiber.Map{
		"blogs": blogs,
	})
}

func Blog(c *fiber.Ctx) error {
	blogId := c.Params("blogId", "")
	if len(blogId) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "")
	}

	b := model.Blog{Id: blogId}
	blog := b.SelectBlog()

	if blog == nil {
		return fiber.ErrNotFound
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"blog": blog,
	})
}

func Likes(c *fiber.Ctx) error {
	blogId := c.Params("blogId", "")
	if len(blogId) == 0 {
		return fiber.ErrBadRequest
	}

	blog := model.Blog{Id: blogId}
	likes := blog.SelectBlogLikes()

	status := fiber.StatusOK
	if len(likes) == 0 {
		status = fiber.StatusNoContent
	}

	return c.Status(status).JSON(fiber.Map{
		"likes": likes,
	})
}

func Comments(c *fiber.Ctx) error {
	blogId := c.Params("blogId", "")
	if len(blogId) == 0 {
		return fiber.ErrBadRequest
	}

	blog := model.Blog{Id: blogId}
	comments := blog.SelectBlogComments()

	status := fiber.StatusOK
	if len(comments) == 0 {
		status = fiber.StatusNoContent
	}

	return c.Status(status).JSON(fiber.Map{
		"comments": comments,
	})
}

func Like(c *fiber.Ctx) error {
	blogId := c.Query("blogId", "")
	userId := c.Query("userId", "")
	if len(blogId) == 0 || len(userId) == 0 {
		return fiber.ErrBadRequest
	}

	return c.SendStatus(fiber.StatusBadRequest)
}

func CreateBlog(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusBadRequest)
}

func CreateComment(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusBadRequest)
}

func UpdateBlog(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusBadRequest)
}

func UpdateComment(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusBadRequest)
}

func DeleteBlog(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusBadRequest)
}

func DeleteComment(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusBadRequest)
}

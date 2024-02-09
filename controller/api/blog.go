package api

import (
	"blog/config"
	"blog/constant"
	"blog/lib/file"
	"blog/model"
	"blog/schema"
	"blog/util"

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

	bl := model.BlogLikes{BlogId: blogId, LikerId: userId}
	bl.ToggleLike()

	return c.SendStatus(fiber.StatusOK)
}

func CreateBlog(c *fiber.Ctx) error {
	body := schema.CreateBlog{}
	if err := c.BodyParser(&body); err != nil {
		panic(err)
	}

	message := util.Validate(body)
	if len(message) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, message)
	}

	form, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}

	converter := file.Converter{Files: form.File["image"]}
	converted, isConverted := converter.Convert()
	if !isConverted {
		return fiber.ErrUnsupportedMediaType
	}

	uploader := file.Uploader{Files: converted}
	uploaded, isUploaded := uploader.Upload()
	if !isUploaded {
		return fiber.ErrInternalServerError
	}

	session, err := config.Store.Get(c)
	if err != nil {
		panic(err)
	}

	sessionUser := session.Get(config.USER).(config.User)
	blog := model.Blog{
		Title:       body.Title,
		Description: body.Description,
		BloggerId:   sessionUser.Id,
	}

	blog.CreateBlog(uploaded...)

	return c.SendStatus(fiber.StatusCreated)
}

func CreateComment(c *fiber.Ctx) error {
	body := schema.CreateComment{}
	if err := c.BodyParser(&body); err != nil {
		panic(err)
	}

	message := util.Validate(body)
	if len(message) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, message)
	}

	blogId := c.Params("blogId", "")
	if len(blogId) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Empty blogId")
	} else if err := util.IsUUID(blogId); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid blogId")
	}

	session, err := config.Store.Get(c)
	if err != nil {
		panic(err)
	}

	sessionUser := session.Get(config.USER).(config.User)
	comment := model.BlogComments{
		BlogId:      blogId,
		CommenterId: sessionUser.Id,
		Comment:     body.Comment,
	}

	comment.CreateComment()

	return c.SendStatus(fiber.StatusCreated)
}

func UpdateBlog(c *fiber.Ctx) error {
	body := schema.UpdateBlog{}
	if err := c.BodyParser(&body); err != nil {
		panic(err)
	}

	message := util.Validate(body)
	if len(message) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, message)
	}

	blogId := c.Params("blogId", "")
	if len(blogId) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Empty blogId")
	} else if err := util.IsUUID(blogId); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid blogId")
	}

	blog := model.Blog{
		Id:          blogId,
		Title:       body.Title,
		Description: body.Description,
	}

	form, err := c.MultipartForm()
	if err != nil {
		panic(err)
	}

	converter := file.Converter{Files: form.File["images"]}
	converted, isConverted := converter.Convert()
	if !isConverted {
		return fiber.ErrUnsupportedMediaType
	}

	uploader := file.Uploader{Files: converted}
	uploaded, isUploaded := uploader.Upload()
	if !isUploaded {
		return fiber.ErrInternalServerError
	}

	blog.UpdateBlog(uploaded...)
	return c.SendStatus(fiber.StatusBadRequest)
}

func UpdateComment(c *fiber.Ctx) error {
	body := schema.UpdateComment{}
	if err := c.BodyParser(&body); err != nil {
		panic(err)
	}

	blogId := c.Params("blogId", "")
	if len(blogId) == 0 {
		return fiber.NewError(fiber.StatusBadRequest,  "Empty blogId")
	} else if err := util.IsUUID(blogId); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid blogId")
	}

	commentId, err := c.ParamsInt("commentId", 0)
	if err != nil {
		panic(err)
	} else if commentId <= 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Empty commentId")
	}

	user := c.Locals(constant.LocalUser).(model.User)
	comment := model.BlogComments{
		Id: commentId,
		BlogId: blogId,
		CommenterId: user.Id,
		Comment: body.Comment,
	}

	comment.UpdateComment()
	return c.SendStatus(fiber.StatusBadRequest)
}

func DeleteBlog(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusBadRequest)
}

func DeleteComment(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusBadRequest)
}

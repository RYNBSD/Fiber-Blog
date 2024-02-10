package api

import (
	"blog/config"
	"blog/lib/file"
	"blog/model"
	"blog/schema"
	"blog/util"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

func Info(c *fiber.Ctx) error {
	userId := c.Params("userId", "")
	if len(userId) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user id")
	}

	user := model.User{Id: userId}
	user.SelectById()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

func Blogs(c *fiber.Ctx) error {
	userId := c.Params("userId", "")
	if len(userId) == 0 {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid user id")
	}

	user := model.User{Id: userId}
	blogs := user.SelectBlogs()

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"blogs": blogs,
	})
}

func Update(c *fiber.Ctx) error {
	body := &schema.Update{}
	if err := c.BodyParser(body); err != nil {
		panic(err)
	}

	message := util.Validate(body)
	if len(message) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, message)
	}

	user := model.User{Email: body.Email}
	if found := user.SelectByEmail(); found {
		return fiber.NewError(fiber.StatusBadRequest, "Email already exists")
	}

	session := config.GetSession(c)
	sessionUser := session.Get(config.USER).(config.User)
	user.Id = sessionUser.Id
	user.Username = body.Username
	user.Picture = ""

	picture, err := c.FormFile("picture")
	if err != nil {
		panic(err)
	}

	if picture != nil {
		converter := file.Converter{Files: []*multipart.FileHeader{picture}}
		converted, isConverted := converter.Convert()
		if !isConverted {
			return fiber.ErrUnsupportedMediaType
		}

		uploader := file.Uploader{Files: converted}
		uploaded, isUploaded := uploader.Upload()
		if !isUploaded {
			return fiber.ErrInternalServerError
		}
		user.Picture = uploaded[0]
	}

	user.Update()
	return c.SendStatus(fiber.StatusOK)
}

func Delete(c *fiber.Ctx) error {
	session := config.GetSession(c)
	sessionUser := session.Get(config.USER).(config.User)
	user := model.User{Id: sessionUser.Id}

	user.Delete()

	session.Set(config.USER, config.User{Id: ""})
	if err := session.Save(); err != nil {
		panic(err)
	}

	return c.SendStatus(fiber.StatusOK)
}

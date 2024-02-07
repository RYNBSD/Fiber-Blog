package api

import (
	"blog/config"
	"blog/lib/file"
	"blog/model"
	"blog/schema"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

func Info(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusBadRequest)
}

func Blogs(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusBadRequest)
}

func Update(c *fiber.Ctx) error {
	body := &schema.Update{}
	if err := c.BodyParser(body); err != nil {
		panic(err)
	}

	user := model.User{Email: body.Email}
	if found := user.SelectByEmail(); found {
		return fiber.NewError(fiber.StatusBadRequest, "Email already exists")
	}

	session, err := config.Store.Get(c)
	if err != nil {
		panic(err)
	}

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
	return c.SendStatus(fiber.StatusBadRequest)
}

func Delete(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusBadRequest)
}

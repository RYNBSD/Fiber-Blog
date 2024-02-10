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

func SignUp(c *fiber.Ctx) error {
	body := schema.SignUp{}
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	message := util.Validate(&body)
	if len(message) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, message)
	}

	user := model.User{
		Username: body.Username,
		Email:    body.Email,
		Password: "",
		Picture:  "",
	}

	if found := user.SelectByEmail(); found {
		return fiber.NewError(fiber.StatusBadRequest, "Email already exists")
	}

	util.EscapeStrings(&body.Username, &body.Password)
	picture, err := c.FormFile("picture")
	if err != nil {
		panic(err)
	}

	convert := file.Converter{Files: []*multipart.FileHeader{picture}}
	converted, isConverted := convert.Convert()
	if !isConverted {
		return fiber.ErrUnsupportedMediaType
	}

	upload := file.Uploader{Files: converted}
	uploaded, isUploaded := upload.Upload()
	if !isUploaded {
		return fiber.ErrInternalServerError
	}
	user.Picture = uploaded[0]

	if user.Password, err = util.HashPassword(body.Password); err != nil {
		panic(err)
	}
	user.Create()

	return c.SendStatus(fiber.StatusCreated)
}

func SignIn(c *fiber.Ctx) error {
	body := schema.SignIn{}
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	message := util.Validate(body)
	if len(message) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, message)
	}

	util.EscapeStrings(&body.Password)
	user := model.User{
		Email: body.Email,
	}

	if found := user.SelectPasswordByEmail(); !found {
		return fiber.ErrNotFound
	}

	if valid := util.ComparePassword(body.Password, user.Password); !valid {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid password")
	}

	if found := user.SelectByEmail(); !found {
		return fiber.ErrNotFound
	}

	session := config.GetSession(c)
	session.Set(config.USER, config.User{Id: user.Id})
	if err := session.Save(); err != nil {
		panic(err)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

func SignOut(c *fiber.Ctx) error {
	session := config.GetSession(c)

	user := config.User{Id: ""}
	access := config.Access{Key: "", Iv: ""}

	session.Set(config.USER, user)
	session.Set(config.ACCESS, access)

	if err := session.Save(); err != nil {
		panic(err)
	}

	return c.SendStatus(fiber.StatusOK)
}

func Me(c *fiber.Ctx) error {
	authorization := c.Get(fiber.HeaderAuthorization, "")
	if len(authorization) == 0 {
		return fiber.ErrUnauthorized
	}

	id, err := util.VerifyJwt(authorization)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	user := model.User{Id: id}
	if found := user.SelectById(); !found {
		return fiber.ErrNotFound
	}

	session := config.GetSession(c)
	session.Set(config.USER, config.User{Id: user.Id})
	if err := session.Save(); err != nil {
		panic(err)
	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": user,
	})
}

func ForgotPassword(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusBadRequest)
}

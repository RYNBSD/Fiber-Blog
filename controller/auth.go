package controller

import (
	"blog/config"
	"blog/lib/file"
	"blog/model"
	"blog/schema"
	"blog/util"
	"encoding/json"
	"mime/multipart"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	body := &schema.SignUp{}
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	message := util.Validate(&body)
	if len(message) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, message)
	}

	// util.EscapeStrings(&body.Username, &body.Password)
	picture, err := c.FormFile("picture")
	if err != nil {
		return err
	}

	convert := file.Converter{Files: []*multipart.FileHeader{picture}}
	converted, isConverted := convert.Convert()
	if !isConverted {
		return fiber.ErrUnsupportedMediaType
	}

	upload := file.Uploader{Files: converted}
	uploaded := upload.Upload()[0]

	user := model.User{
		Username: body.Username,
		Email:    body.Email,
		Password: "",
		Picture:  uploaded,
	}

	if user.Password, err = util.HashPassword(body.Password); err != nil {
		panic(err)
	}
	user.Create()

	return c.SendStatus(fiber.StatusCreated)
}

func SignIn(c *fiber.Ctx) error {
	body := &schema.SignIn{}
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	message := util.Validate(body)
	if len(message) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, message)
	}

	// util.EscapeStrings(&body.Password)

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

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

func SignOut(c *fiber.Ctx) error {
	session, err := config.Store.Get(c)
	if err != nil {
		panic(err)
	}

	user, err := json.Marshal(config.User{Id: ""})
	if err != nil {
		panic(err)
	}

	access, err := json.Marshal(config.Access{Key: "", Iv: ""})
	if err != nil {
		panic(err)
	}

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

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"user": user,
	})
}

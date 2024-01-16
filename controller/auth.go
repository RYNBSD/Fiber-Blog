package controller

import (
	"blog/schema"
	"blog/util"
	"io"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	var body *schema.SignUp
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	message := util.Validate(body)
	if len(message) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, message)
	}

	util.EscapeStrings(&body.Username, &body.Password)

	picture, err := c.FormFile("picture")
	if err != nil {
		return err
	}

	return c.SendStatus(fiber.StatusCreated)
}

func SignIn(c *fiber.Ctx) error {
	var body *schema.SignIn
	if err := c.BodyParser(&body); err != nil {
		return err
	}

	message := util.Validate(body)
	if len(message) > 0 {
		return fiber.NewError(fiber.StatusBadRequest, message)
	}

	util.EscapeStrings(&body.Password)

	return c.SendStatus(fiber.StatusOK)
}

func SignOut(c *fiber.Ctx) error {
	return c.SendStatus(fiber.StatusOK)
}

func Me(c *fiber.Ctx) error {
	authorization := c.Get(fiber.HeaderAuthorization, "")
	if len(authorization) == 0 {
		return fiber.ErrUnauthorized
	}

	_, err := util.VerifyJwt(authorization)
	if err != nil {
		return fiber.ErrUnauthorized
	}

	// TODO: add function to verify if user exists

	return c.SendStatus(fiber.StatusCreated)
}

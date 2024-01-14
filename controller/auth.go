package controller

import (
	"blog/schema"
	"blog/util"

	"github.com/gofiber/fiber/v2"
)

func SignUp(c *fiber.Ctx) error {
	var body *schema.SignUp
	if err := c.BodyParser(&body); err != nil {

	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
	})
}

func SignIn(c *fiber.Ctx) error {
	var body *schema.SignIn
	if err := c.BodyParser(&body); err != nil {

	}

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
	}) 
}

func SignOut(c *fiber.Ctx) error {
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
	})
}

func Me(c *fiber.Ctx) error {
	authorization := c.Get("authorization", "")
	if len(authorization) == 0 {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": "Empty authorization key",
		})
	}

	_, err := util.VerifyJwt(authorization)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"success": false,
			"message": err.Error(),
		})
	}

	// TODO: add function to verify if user exists

	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
	})
}

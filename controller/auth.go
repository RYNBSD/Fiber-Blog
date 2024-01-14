package controller

import (
	"blog/schema"

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
	return c.Status(fiber.StatusCreated).JSON(fiber.Map{
		"success": true,
	})
}

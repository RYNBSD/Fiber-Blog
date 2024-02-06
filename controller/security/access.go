package security

import "github.com/gofiber/fiber/v2"

func AccessEmail(c *fiber.Ctx) error {

	return c.SendStatus(fiber.StatusBadRequest)
}
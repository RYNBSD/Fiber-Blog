package middleware

import (
	"blog/config"
	"blog/constant"
	"blog/model"
	"blog/util"

	"github.com/gofiber/fiber/v2"
)

func HasUserRegistered(c *fiber.Ctx) error {
	userId := c.Get(constant.HttpHeaderUserId, "")
	if len(userId) == 0 {
		return fiber.NewError(fiber.StatusUnauthorized, "Empty user id (header)")
	} else if err := util.IsUUID(userId); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user id (uuid)")
	}

	session, err := config.Store.Get(c)
	if err != nil {
		panic(err)
	}

	sessionUser, ok := session.Get(config.USER).(config.User)
	if !ok {
		panic("Invalid user session")
	} else if err := util.IsUUID(sessionUser.Id); err != nil {
		panic(err)
	}

	if (userId != sessionUser.Id) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user id")
	}

	user := model.User{Id: userId}
	if found := user.SelectById(); !found {

		session.Set(config.USER, config.User{Id: ""})
		if err := session.Save(); err != nil {
			panic(err)
		}

		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	c.Locals(constant.LocalUser, user)
	return c.Next()
}

func HasUserUnregistered(c *fiber.Ctx) error {
	userId := c.Get(constant.HttpHeaderUserId, "")
	if len(userId) == 0 {
		return c.Next()
	} else if err := util.IsUUID(userId); err != nil {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user id (uuid)")
	}

	session, err := config.Store.Get(c)
	if err != nil {
		panic(err)
	}

	sessionUser, ok := session.Get(config.USER).(config.User)
	if !ok || len(sessionUser.Id) == 0 {
		return c.Next()
	} else if err := util.IsUUID(sessionUser.Id); err != nil {
		panic(err)
	}

	if (userId != sessionUser.Id) {
		return fiber.NewError(fiber.StatusUnauthorized, "Invalid user id")
	}

	user := model.User{Id: userId}
	if found := user.SelectById(); !found {

		session.Set(config.USER, config.User{Id: ""})
		if err := session.Save(); err != nil {
			panic(err)
		}

		return fiber.NewError(fiber.StatusNotFound, "User not found")
	}

	c.Locals(constant.LocalUser, user)
	return c.Next()
}
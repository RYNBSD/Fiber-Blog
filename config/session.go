package config

import (
	"encoding/gob"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
)

var store *session.Store = session.New()

func GetSession(c *fiber.Ctx) *session.Session {
	session, err := store.Get(c)
	if err != nil {
		panic(err)
	}

	return session
}

func InitGob() {
	gob.Register(User{})
	gob.Register(Access{})
}

const (
	USER   = "user"
	ACCESS = "access"
)

type User struct {
	Id string
}

type Access struct {
	Key string
	Iv  string
}

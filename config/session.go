package config

import "github.com/gofiber/fiber/v2/middleware/session"

var Store *session.Store = session.New()

const (
	USER = "user"
	ACCESS = "access"
)

type User struct {
	Id string
}

type Access struct {
	Key string
	Iv string
}
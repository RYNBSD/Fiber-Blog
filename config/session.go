package config

import "github.com/gofiber/fiber/v2/middleware/session"

var Store *session.Store = session.New()

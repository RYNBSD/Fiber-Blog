package util

import (
	"blog/constant"
	"fmt"
	"html"
	"os"
	"path"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

func PublicDir() string {
	root := rootDir()
	return path.Join(root, constant.PUBLIC)
}

func EscapeStrings(strings ...*string) {
	for _, str := range strings {
		*str = html.EscapeString(*str)
	}
}

func Validate(data any) string {
	validate := validator.New()
	errs := validate.Struct(data)
	message := ""

	if errs != nil {
		for _, err := range errs.(validator.ValidationErrors) {
			message += fmt.Sprintf("%v: %v;", err.Field(), err.Tag())
		}
	}

	return message
}

func UUIDv4() string {
	v4, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	uuid := v4.String()
	if len(uuid) == 0 {
		panic("Empty uuid v4")
	}

	return uuid
}

func rootDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}


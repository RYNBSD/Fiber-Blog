package util

import (
	"blog/constant"
	"fmt"
	"html"
	"os"
	"path"

	"github.com/go-playground/validator/v10"
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

func rootDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}


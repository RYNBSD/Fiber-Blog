package util

import (
	"blog/constant"
	"html"
	"os"
	"path"
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

func rootDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}


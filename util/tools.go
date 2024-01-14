package util

import (
	"os"
	"path"
)

func RootDir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}

const PUBLIC = "public"

func PublicDir() string {
	root := RootDir()
	return path.Join(root, PUBLIC)
}
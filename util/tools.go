package util

import (
	"blog/constant"
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

func PublicDir() string {
	root := RootDir()
	return path.Join(root, constant.PUBLIC)
}
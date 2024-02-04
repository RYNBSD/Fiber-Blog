package file

import (
	"blog/constant"
	"blog/util"
	"fmt"
	"os"
	"path"
	"sync"
	"time"

	"github.com/google/uuid"
)

type IUploader interface {
	Upload() []string
	Remove(...string)

	uniqueFileName() string
}

type Uploader struct {
	Files [][]byte
}

func (u *Uploader) Upload() []string {
	publicDir := util.PublicDir()
	uploaded := make([]string, 0)
	wg := sync.WaitGroup{}

	for _, file := range u.Files {
		wg.Add(1)

		go func(file []byte) {
			defer wg.Done()

			name := u.uniqueFileName()
			fullPath := path.Join(publicDir, name)

			open, err := os.Create(fullPath)
			if err != nil {
				panic(err)
			}

			_, err = open.Write(file)
			if err != nil {
				panic(err)
			}

			defer open.Close()
			uploaded = append(uploaded, name)
		}(file)
	}

	wg.Wait()
	return uploaded
}

func (u *Uploader) Remove(paths ...string) {
	wg := sync.WaitGroup{}

	for _, path := range paths {
		wg.Add(1)

		go func(path string) {
			defer wg.Done()

			if err := os.Remove(path); err != nil {
				panic(err)
			}
		}(path)
	}
	wg.Wait()
}

func (u *Uploader) uniqueFileName() string {
	uuid, err := uuid.NewRandom()
	if err != nil {
		panic(err)
	}

	second := fmt.Sprintf("%v", time.Now().Second())
	return second + "_" + uuid.String() + "." + constant.JPEG
}

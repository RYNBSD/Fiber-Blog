package file

import (
	"io"
	"mime/multipart"
	"sync"

	"github.com/h2non/bimg"
	"github.com/h2non/filetype"
)

type IConverter interface {
	Convert() []byte

	verifyImages()
	// Converted Images, Number of Converted Images
	toWebp() ([][]byte, bool)
}

type Converter struct {
	Files []*multipart.FileHeader
	To    string

	filesContent [][]byte
}

func (c *Converter) verifyImages() {
	filteredFiles := make([][]byte, 0)
	wg := sync.WaitGroup{}

	for _, file := range c.Files {
		wg.Add(1)
		go func(file *multipart.FileHeader) {
			defer wg.Done()

			open, err := file.Open()
			if err != nil {
				panic(err)
			}

			read, err := io.ReadAll(open)
			if err != nil {
				panic(err)
			}

			if filetype.IsImage(read) {
				filteredFiles = append(filteredFiles, read)
			}
		}(file)
	}
	wg.Wait()
	c.filesContent = filteredFiles
}

func (c *Converter) toWebp() [][]byte {
	wg := sync.WaitGroup{}
	WebPs := make([][]byte, 0)

	for _, file := range c.filesContent {
		wg.Add(1)
		go func(file []byte) {
			defer wg.Done()

			webp, err := bimg.NewImage(file).Convert(bimg.WEBP)
			if err != nil {
				panic(err)
			}
			WebPs = append(WebPs, webp)
		}(file)
	}
	wg.Wait()

	return WebPs
}

//
func (c *Converter) Convert() ([][]byte, bool) {
	c.verifyImages()
	webps := c.toWebp()
	return webps, len(webps) > 0
}

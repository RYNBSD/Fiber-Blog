package file

import (
	"bytes"
	"image"
	"image/jpeg"
	"io"
	"mime/multipart"
	"sync"

	"github.com/h2non/filetype"
)

type IConverter interface {
	Convert() ([][]byte, bool)

	verifyImages()
	toJpeg() [][]byte
}

type Converter struct {
	Files []*multipart.FileHeader

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
			defer open.Close()

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

func (c *Converter) toJpeg() [][]byte {
	wg := sync.WaitGroup{}
	JPEGs := make([][]byte, 0)

	for _, file := range c.filesContent {
		wg.Add(1)

		go func(file []byte) {
			defer wg.Done()

			img, _, err := image.Decode(bytes.NewReader(file))
			if err != nil {
				panic(err)
			}

			buffer := bytes.Buffer{}
			options := jpeg.Options{
				Quality: 100,
			}

			if err := jpeg.Encode(&buffer, img, &options); err != nil {
				panic(err)
			}

			jpeg := buffer.Bytes()
			JPEGs = append(JPEGs, jpeg)
		}(file)
	}
	wg.Wait()

	return JPEGs
}

// Converted Images, Number of Converted Images
func (c *Converter) Convert() ([][]byte, bool) {
	c.verifyImages()
	jpegs := c.toJpeg()
	return jpegs, len(jpegs) > 0
}

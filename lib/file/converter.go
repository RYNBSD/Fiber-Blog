package file

import (
	"io"
	"mime/multipart"
	"sync"

	"github.com/h2non/bimg"
	"github.com/h2non/filetype"
)

type IConverter interface {
	Convert() ([][]byte, bool)

	verifyImages()
	toWebp() [][]byte
}

type Converter struct {
	Files        []*multipart.FileHeader
	filesContent [][]byte
}

func (c *Converter) verifyImages() {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	filteredFiles := make([][]byte, 0)

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
				mutex.Lock()
				filteredFiles = append(filteredFiles, read)
				mutex.Unlock()
			}
		}(file)
	}
	wg.Wait()
	c.filesContent = filteredFiles
}

func (c *Converter) toWebp() [][]byte {
	wg := sync.WaitGroup{}
	mutex := sync.Mutex{}
	WebPs := make([][]byte, 0)

	for _, file := range c.filesContent {
		wg.Add(1)
		go func(file []byte) {
			defer wg.Done()

			opt := bimg.Options{
				Width:     1280,
				Height:    720,
				Quality:   100,
				Interlace: true,
				Enlarge:   true,
				Embed:     true,
				Force:     true,
				Crop:      true,
				Gravity:   bimg.GravitySmart,
				Type:      bimg.WEBP,
			}

			webp, err := bimg.NewImage(file).Process(opt)
			if err != nil {
				panic(err)
			}

			mutex.Lock()
			WebPs = append(WebPs, webp)
			mutex.Unlock()
		}(file)
	}
	wg.Wait()

	return WebPs
}

// Converted Images, Number of Converted Images
func (c *Converter) Convert() ([][]byte, bool) {
	c.verifyImages()
	webps := c.toWebp()
	return webps, len(webps) > 0
}

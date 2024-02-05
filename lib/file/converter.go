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
	Files []*multipart.FileHeader
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

			options := bimg.Options{
				Width:        1280,            // Set the width of the output image
				Height:       720,            // Set the height of the output image
				Quality:      100,             // Set the quality of the output image (0-100)
				Interlace:    true,           // Enable progressive (interlaced) rendering
				Enlarge:      true,           // Allow enlarging images (by default, bimg prevents upscaling)
				Embed:        true,           // Embed ICC profiles and comments
				Gravity:      bimg.GravitySmart, // Set the gravity for resizing (e.g., bimg.GravityNorthWest)
				Type:         bimg.WEBP,      // Set the output image format (e.g., bimg.WEBP, bimg.PNG)
			}

			webp, err := bimg.NewImage(file).Process(options)
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
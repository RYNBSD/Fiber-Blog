package model

import (
	"blog/util"
	"os"
	"path"
	"sync"
)

func createImages(blogId string, images ...string) {
	if images != nil {
		wg := sync.WaitGroup{}

		const sql = `DELETE FROM "blogImages" WHERE "blogId"=?`
		if _, err := DB.Exec(sql, blogId); err != nil {
			panic(err)
		}

		for _, image := range images {
			wg.Add(1)

			go func(image string) {
				defer wg.Done()
				const sql = `INSERT INTO "blogImages" (image, "blogId") VALUES (?, ?)`

				if _, err := DB.Exec(sql, image, blogId); err != nil {
					panic(err)
				}
			}(image)
		}
		wg.Wait()
	}
}

func deleteImages(images ...string) {
	if images != nil {
		wg := sync.WaitGroup{}

		for _, image := range images {
			wg.Add(1)

			go func(image string) {
				defer wg.Done()
				imagePath := path.Join(util.PublicDir(), image)

				if err := os.Remove(imagePath); err != nil {
					panic(err)
				}
			}(image)
		}
		wg.Wait()
	}
}

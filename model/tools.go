package model

import "sync"

func createImages(blogId string, images ...string) {
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

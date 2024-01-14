package model

import "sync"

func CreateBlog(blog Blog) {
	Connect()
	sql := `INSERT INTO blog (title, description, "bloggerId") VALUES (?, ?, ?)`

	if _, err := DB.Exec(sql, blog.Title, blog.Description, blog.BloggerId); err != nil {
		panic(err)
	}
}

func UpdateBlog(blog Blog, images ...string) {
	Connect()

	if images != nil {
		wg := sync.WaitGroup{}

		sql := `DELETE FROM "blogImages" WHERE "blogId"=?`
		if _, err := DB.Exec(sql, blog.Id); err != nil {
			panic(err)
		}

		for _, image := range images {
			wg.Add(1)

			go func(image string) {
				defer wg.Done()
				sql := `INSERT INTO "blogImages" (image, "blogId") VALUES (?, ?)`

				if _, err := DB.Exec(sql, image, blog.Id); err != nil {
					panic(err)
				}
			}(image)
		}
		wg.Done()
	}

	sql := ``
	if _, err := DB.Exec(sql); err != nil {
		panic(err)
	}
}

func DeleteBlog(id string) {
	Connect()
	sql := `DELETE FROM blog WHERE id=?`

	if _, err := DB.Exec(sql, id); err != nil {
		panic(err)
	}
}

func SelectBlogComments() {

}

func SelectBlogLikes() {

}

func NewBlogLike(like BlogLikes) {
	Connect()
	sql := `INSERT INTO "blogLikes" ("likerId", "blogId") VALUES (?, ?)`

	if _, err := DB.Exec(sql, like.LikerId, like.BlogId); err != nil {
		panic(err)
	}
}

func NewBlogComment(comment BlogComments) {
	Connect()
	sql := `INSERT INTO "blogComments" ("commenterId", "blogId", comment) VALUES (?, ?, ?)`

	if _, err := DB.Exec(sql, comment.CommenterId, comment.BlogId, comment.Comment); err != nil {
		panic(err)
	}
}

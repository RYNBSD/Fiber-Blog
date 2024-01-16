package model

func CreateBlog(blog Blog, images ...string) {
	Connect()
	createImages(blog.Id, images...)
	const sql = `INSERT INTO blog (title, description, "bloggerId") VALUES (?, ?, ?)`

	if _, err := DB.Exec(sql, blog.Title, blog.Description, blog.BloggerId); err != nil {
		panic(err)
	}
}

func UpdateBlog(blog Blog, images ...string) {
	Connect()
	createImages(blog.Id, images...)
	const sql = `UPDATE blog`

	if _, err := DB.Exec(sql); err != nil {
		panic(err)
	}
}

func DeleteBlog(id string) {
	Connect()

	const sql1 = `SELECT image from "blogImages" WHERE id=?`
	if rows, err := DB.Query(sql1, id); err != nil {
		panic(err)
	} else {
		defer rows.Close()
		images := make([]string, 0)

		for rows.Next() {
			image := ""
			if err := rows.Scan(&image); err != nil {
				panic(err)
			}
			images = append(images, image)
		}

		if err := rows.Err(); err != nil {
			panic(err)
		}
		deleteImages(images...)
	}

	const sql = `DELETE FROM blog WHERE id=?`
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
	const sql = `INSERT INTO "blogLikes" ("likerId", "blogId") VALUES (?, ?)`

	if _, err := DB.Exec(sql, like.LikerId, like.BlogId); err != nil {
		panic(err)
	}
}

func NewBlogComment(comment BlogComments) {
	Connect()
	const sql = `INSERT INTO "blogComments" ("commenterId", "blogId", comment) VALUES (?, ?, ?)`

	if _, err := DB.Exec(sql, comment.CommenterId, comment.BlogId, comment.Comment); err != nil {
		panic(err)
	}
}

func RemoveBlogLike() {}

func RemoveBlogComment() {}

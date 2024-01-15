package model

func CreateBlog(blog Blog, images ...string) {
	Connect()
	if images != nil {
		createImages(blog.Id, images...)
	}

	const sql = `INSERT INTO blog (title, description, "bloggerId") VALUES (?, ?, ?)`

	if _, err := DB.Exec(sql, blog.Title, blog.Description, blog.BloggerId); err != nil {
		panic(err)
	}
}

func UpdateBlog(blog Blog, images ...string) {
	Connect()

	if images != nil {
		createImages(blog.Id, images...)
	}

	const sql = ``
	if _, err := DB.Exec(sql); err != nil {
		panic(err)
	}
}

func DeleteBlog(id string) {
	Connect()
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

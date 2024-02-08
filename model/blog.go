package model

import (
	"blog/types"
	"blog/util"
	"database/sql"
)

func (b *Blog) CreateBlog(images ...string) {
	Connect()
	id := util.UUIDv4()

	const sql = `INSERT INTO "blog" ("id", "title", "description", "bloggerId") VALUES ($1, $2, $3)`
	if _, err := DB.Exec(sql, id, b.Title, b.Description, b.BloggerId); err != nil {
		panic(err)
	}

	createImages(id, images...)
}

func (b *Blog) UpdateBlog(images ...string) {
	Connect()
	createImages(b.Id, images...)
	const sql = `UPDATE blog`

	if _, err := DB.Exec(sql); err != nil {
		panic(err)
	}
}

func (b *Blog) DeleteBlog() {
	Connect()

	const imagesSql = `SELECT image from "blogImages" WHERE "blogId"=$1`
	rows, err := DB.Query(imagesSql, b.Id)

	if err != nil {
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

	const sql = `DELETE FROM "blog" WHERE "id"=$1`
	if _, err := DB.Exec(sql, b.Id); err != nil {
		panic(err)
	}
}

func (b *Blog) SelectBlogs() []types.Map {
	Connect()

	const query = `
	SELECT b.id, b.title, b.description, ARRAY_AGG(bi.image) AS images, u.id AS "bloggerId", u.username, u.picture,
	COUNT(bl.id) AS likes, COUNT(bc.id) AS comments
	FROM blog b
	INNER JOIN "user" u ON u."id" = b."bloggerId"
	LEFT JOIN "blogImages" bi ON bi."blogId" = b."id"
	LEFT JOIN "blogLikes" bl ON bl."blogId" = b."id"
	LEFT JOIN "blogComments" bc ON bc."blogId" = b."id"
	GROUP BY b.id, b.title, b.description, u.id, u.username, u.picture
	`

	rows, err := DB.Query(query)
	if err != nil {
		panic(err)
	}

	blogs := []types.Map{}
	scanUnknownColumns(rows, &blogs)
	return blogs
}

func (b *Blog) SelectBlog() types.Map {
	Connect()

	const query = `
	SELECT b.id, b.title, b.description, ARRAY_AGG(bi.image) AS images, u.id AS "bloggerId", u.username, u.picture, COUNT(bl.id) AS likes, COUNT(bc.id) AS comments FROM blog b
	INNER JOIN "user" u ON u."id" = b."bloggerId"
	LEFT JOIN "blogLikes" bl ON bl."blogId" = b."id"
	LEFT JOIN "blogComments" bc ON bc."blogId" = b."id"
	LEFT JOIN "blogImages" bi ON bi."blogId" = b."id"
	WHERE b.id = $1
	GROUP BY b.id, b.title, b.description, u.id, u.username, u.picture
	`

	rows, err := DB.Query(query)
	if err != nil {
		panic(err)
	}

	blogs := []types.Map{}
	scanUnknownColumns(rows, &blogs)

	if len(blogs) == 0 {
		return nil
	}
	return blogs[0]
}

func (b *Blog) SelectBlogLikes() []types.Map {
	Connect()
	const query = `
	SELECT u.username, u.picture FROM "blogLikes" bl
	INNER JOIN "user" u ON u.id = bl."likerId"
	WHERE bl."blogId"= $1
	`

	rows, err := DB.Query(query, b.Id)
	if err != nil {
		panic(err)
	}

	likes := []types.Map{}
	scanUnknownColumns(rows, &likes)

	return likes
}

func (b *Blog) SelectBlogComments() []types.Map {
	Connect()
	const query = `
	SELECT u.username, u.picture, bc.comment FROM "blogComments" bc
	INNER JOIN "user" u ON u.id = bc."commenterId"
	WHERE bc."blogId" = $1
	`

	rows, err := DB.Query(query, b.Id)
	if err != nil {
		panic(err)
	}

	comments := []types.Map{}
	scanUnknownColumns(rows, &comments)

	return comments
}

func (l *BlogLikes) ToggleLike() bool {
	Connect()

	found := false

	const find = `SELECT "id" FROM "blogLikes" WHERE "blogId"=$1 AND "userId"=$2 LIMIT 1`
	_, err := DB.Query(find, l.BlogId, l.LikerId)
	switch err {
	case sql.ErrNoRows:
		found = false
	case nil:
		found = true
	default:
		panic(err)
	}

	if found {
		const delete = `DELETE FROM "blogLikes" WHERE "blogId"=$1 AND "userId"=$2`
		_, err := DB.Exec(delete, l.BlogId, l.LikerId)
		if err != nil {
			panic(err)
		}
	} else {
		const create = `INSERT INTO "blogLikes"("blogId", "userId") VALUES ($1, $2)`
		_, err := DB.Exec(create, l.BlogId, l.LikerId)
		if err != nil {
			panic(err)
		}
	}

	return !found
}

func (c *BlogComments) CreateComment() {
	Connect()
	const sql = `INSERT INTO "blogComments" ("commenterId", "blogId", "comment") VALUES ($1, $2, $3)`

	if _, err := DB.Exec(sql, c.CommenterId, c.BlogId, c.Comment); err != nil {
		panic(err)
	}
}

func (c *BlogComments) UpdateComment() {

}

func (c *BlogComments) DeleteComment() {}

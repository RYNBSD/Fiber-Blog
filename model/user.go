package model

import (
	"blog/types"
	"blog/util"
	"database/sql"
	"sync"
)

func (u *User) Create() {
	Connect()
	uuid := util.UUIDv4()

	const query = `INSERT INTO "user"("id", "username", "email", "password", "picture") VALUES ($1, $2, $3, $4, $5)`

	if _, err := DB.Exec(query, uuid, u.Username, u.Email, u.Password, u.Picture); err != nil {
		panic(err)
	}
}

func (u *User) Update() {
	Connect()

	if len(u.Picture) > 0 {
		const query = `UPDATE "user" SET "username"=$1, "email"=$2, "picture"=$3, "updatedAt"=NOW() WHERE "id"=$4`
		if _, err := DB.Exec(query, u.Username, u.Email, u.Picture, u.Id); err != nil {
			panic(err)
		}
	} else {
		const query = `UPDATE "user" SET "username"=$1, "email"=$2, "updatedAt"=NOW() WHERE "id"=$3`
		if _, err := DB.Exec(query, u.Username, u.Email, u.Id); err != nil {
			panic(err)
		}
	}
}

func (u *User) Delete() {
	Connect()

	// Delete user picture from public
	picture := ""
	row := DB.QueryRow("SELECT \"picture\" FROM \"user\" WHERE \"id\"=$1 LIMIT 1", u.Id)

	if err := row.Err(); err != nil {
		panic(err)
	}

	if err := row.Scan(&picture); err != nil {
		panic(err)
	}
	deleteImages(picture)

	// Get blog IDs to delete there images
	IDs := make([]string, 0)
	rows, err := DB.Query("SELECT \"id\" FROM \"blog\" WHERE \"bloggerId\"=$1", u.Id)

	if err != nil {
		panic(err)
	} else {
		defer rows.Close()
		id := ""

		for rows.Next() {
			if err := rows.Scan(&id); err != nil {
				panic(err)
			}
			IDs = append(IDs, id)
		}

		if err := rows.Err(); err != nil {
			panic(err)
		}
	}

	// DELETE Blogs With Images
	wg := sync.WaitGroup{}
	for _, id := range IDs {
		wg.Add(1)

		go func(id string) {
			defer wg.Done()
			blog := Blog{Id: id}
			blog.DeleteBlog()
		}(id)
	}
	wg.Wait()

	if _, err := DB.Exec("DELETE FROM \"user\" WHERE \"id\"=$1", u.Id); err != nil {
		panic(err)
	}
}

// This is the default for select user
func (u *User) SelectById() bool {
	Connect()

	if len(u.Id) == 0 {
		panic("Unprovided id to select user")
	} else if err := util.IsUUID(u.Id); err != nil {
		panic(err)
	}

	const query = `
		SELECT "username", "email", "picture"
		FROM "user"
		WHERE "id"=$1
		LIMIT 1
	`

	row := DB.QueryRow(query, u.Id)
	err := row.Err()

	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err)
	}

	if err := row.Scan(&u.Username, &u.Email, &u.Picture); err == sql.ErrNoRows {
		return false
	}
	u.Password = ""
	return true
}

func (u *User) SelectByEmail() bool {
	Connect()

	if len(u.Email) == 0 {
		panic("Unprovided email to select user")
	} else if err := util.IsEmail(u.Email); err != nil {
		panic(err)
	}

	const query = `
		SELECT "id", "username", "picture"
		FROM "user"
		WHERE "email"=$1
		LIMIT 1
	`

	row := DB.QueryRow(query, u.Email)
	err := row.Err()

	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err)
	}

	if err := row.Scan(&u.Id, &u.Username, &u.Picture); err == sql.ErrNoRows {
		return false
	}
	u.Password = ""
	return true
}

func (u *User) SelectBlogs() []types.Map {
	Connect()

	const query = `
	SELECT b.id, b.title, b.description, ARRAY_AGG(bi.image) AS images, u.id AS "bloggerId", u.username, u.picture, COUNT(bl.id) AS likes, COUNT(bc.id) AS comments FROM blog b
	LEFT JOIN "user" u ON u."id" = b."bloggerId"
	LEFT JOIN "blogLikes" bl ON bl."blogId" = b."id"
	LEFT JOIN "blogComments" bc ON bc."blogId" = b."id"
	LEFT JOIN "blogImages" bi ON bi."blogId" = b."id"
	WHERE u.id = $1
	GROUP BY b.id, b.title, b.description, u.id, u.username, u.picture
	`

	rows, err := DB.Query(query, u.Id)
	if err != nil {
		panic(err)
	}

	blogs := []types.Map{}
	scanUnknownColumns(rows, &blogs)
	return blogs
}

func (u *User) SelectPasswordById() bool {
	Connect()

	if err := util.IsUUID(u.Id); err != nil {
		panic(err)
	}

	const query = `
		SELECT "password"
		FROM "user"
		WHERE "id"=$1
		LIMIT 1
	`
	row := DB.QueryRow(query, u.Id)
	err := row.Err()

	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err)
	}

	row.Scan(&u.Password)
	return true
}

func (u *User) SelectPasswordByEmail() bool {
	Connect()

	if err := util.IsEmail(u.Email); err != nil {
		panic(err)
	}

	const query = `
		SELECT "password"
		FROM "user"
		WHERE "email"=$1
		LIMIT 1
	`
	row := DB.QueryRow(query, u.Email)
	err := row.Err()

	switch {
	case err == sql.ErrNoRows:
		return false
	case err != nil:
		panic(err)
	}

	if err := row.Scan(&u.Password); err == sql.ErrNoRows {
		return false
	}
	return true
}

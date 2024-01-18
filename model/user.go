package model

import (
	"blog/types"
	"blog/util"
	"strings"
	"sync"
)

func (u *User) CreateUser() {
	Connect()
	const sql = `INSERT INTO user(username, email, password, picture) VALUES (?, ?, ?, ?)`

	if _, err := DB.Exec(sql, u.Username, u.Email, u.Password, u.Picture); err != nil {
		panic(err)
	}
}

func (u *User) UpdateUser() {
	Connect()

	if len(u.Picture) > 0 {
		const sql = `UPDATE user SET username=?, email=?, password=?, picture=? WHERE id=?`
		if _, err := DB.Exec(sql, u.Username, u.Email, u.Password, u.Picture, u.Id); err != nil {
			panic(err)
		}
	} else {
		const sql = `UPDATE user SET username=?, email=?, password=? WHERE id=?`
		if _, err := DB.Exec(sql, u.Username, u.Email, u.Password, u.Id); err != nil {
			panic(err)
		}
	}
}

func (u *User) DeleteUser() {
	Connect()
	id := u.Id

	// Delete user picture from public
	picture := ""
	row := DB.QueryRow("SELECT picture FROM user WHERE id=? LIMIT 1", id)
	if err := row.Scan(&picture); err != nil {
		panic(err)
	}
	deleteImages(picture)

	// Get blog IDs to delete there images
	IDs := make([]string, 0)
	if rows, err := DB.Query("SELECT id FROM blog WHERE \"bloggerId\"=?", id); err != nil {
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

	if _, err := DB.Exec("DELETE FROM user WHERE id=?", id); err != nil {
		panic(err)
	}
}

func (u *User) VerifyUser(scan bool) bool {
	Connect()
	sql := `
		SELECT
		CASE WHEN COUNT($) > 0 THEN 'true' ELSE 'false' END AS found FROM user
		WHERE $=?
		LIMIT 1
	`
	by := ""
	found := false

	if len(u.Id) > 0 {

		by = u.Id
		if err := util.IsUUID(by); err != nil {
			panic(err)
		}

		sql = strings.Replace(sql, "$", "id", 2)
	} else if len(u.Email) > 0 {

		by = u.Email
		if err := util.IsEmail(by); err != nil {
			panic(err)
		}

		sql = strings.Replace(sql, "$", "email", 2)
	} else {
		panic("Unprovided id or email to verify if user exist")
	}

	row := DB.QueryRow(sql, by)
	if err := row.Err(); err != nil {
		panic(err)
	}
	row.Scan(&found)

	if found {
		u.SelectUser()
	}
	return found
}

func (u *User) SelectUser() {
	Connect()

	if len(u.Id) == 0 {
		panic("Unprovided user id to select user")
	} else if err := util.IsUUID(u.Id); err != nil {
		panic(err)
	}
	id := u.Id

	const sql = `
		SELECT username, email, picture
		FROM user
		WHERE id=?
		LIMIT 1
	`

	row := DB.QueryRow(sql, id)
	if err := row.Err(); err != nil {
		panic(err)
	}

	row.Scan(&u.Username, &u.Email, &u.Picture)
}

func (u *User) ProfileUser() []types.Map {
	Connect()

	u.SelectUser()
	id := u.Id
	const sql = ``

	rows, err := DB.Query(sql, id)
	if err != nil {
		panic(err)
	}

	defer rows.Close()
	blogs := make([]types.Map, 0)
	scanUnknownColumns(rows, &blogs)

	if err := rows.Err(); err != nil {
		panic(err)
	}
	return blogs
}

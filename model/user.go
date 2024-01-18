package model

import (
	"blog/types"
	"reflect"
	"strings"
	"sync"

	"github.com/google/uuid"
)

func (user User) CreateUser() {
	Connect()
	const sql = `INSERT INTO user(username, email, password, picture) VALUES (?, ?, ?, ?)`

	if _, err := DB.Exec(sql, user.Username, user.Email, user.Password, user.Picture); err != nil {
		panic(err)
	}
}

func (user User) UpdateUser() {
	Connect()

	if len(user.Picture) > 0 {
		const sql = `UPDATE user SET username=?, email=?, password=?, picture=? WHERE id=?`
		if _, err := DB.Exec(sql, user.Username, user.Email, user.Password, user.Picture, user.Id); err != nil {
			panic(err)
		}
	} else {
		const sql = `UPDATE user SET username=?, email=?, password=? WHERE id=?`
		if _, err := DB.Exec(sql, user.Username, user.Email, user.Password, user.Id); err != nil {
			panic(err)
		}
	}
}

func (user User) DeleteUser() {
	Connect()
	id := user.Id

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

func (user User) VerifyUser() bool {
	Connect()
	sql := `
		SELECT
		CASE WHEN COUNT($) > 0 THEN 'true' ELSE 'false' END AS found FROM user
		WHERE $=?
		LIMIT 1
	`
	by := ""
	found := false

	if len(user.Id) > 0 {
		by = user.Id
		sql = strings.Replace(sql, "$", "id", 2)
	} else if len(user.Email) > 0 {
		by = user.Email
		sql = strings.Replace(sql, "$", "email", 2)
	} else {
		panic("Unprovided id or email to verify if user exist")
	}

	row := DB.QueryRow(sql, by)
	if err := row.Err(); err != nil {
		panic(err)
	}
	row.Scan(&found)
	return found
}

func (user *User) SelectUser() {
	Connect()

	if len(user.Id) == 0 {
		panic("Unprovided user id to select user")
	} else if _, err := uuid.Parse(user.Id); err != nil {
		panic(err)
	}
	id := user.Id

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

	row.Scan(&user.Username, &user.Email, &user.Picture)
}

func (user User) ProfileUser() []types.Map {
	Connect()

	user.SelectUser()
	id := user.Id
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

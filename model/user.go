package model

import (
	"sync"
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

func (user User) SelectUser() {
	Connect()
	id := user.Id
	const sql = ``

	rows, err := DB.Query(sql, id)
	if err != nil {
		panic(err)
	} else {
		defer rows.Close()
	}
}

func (user User) ProfileUser() {
	Connect()
	id := user.Id
	const sql = ``

	rows, err := DB.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}

package model

import "sync"

func CreateUser(user User) {
	Connect()
	const sql = `INSERT INTO user(username, email, password, picture) VALUES (?, ?, ?, ?)`

	if _, err := DB.Exec(sql, user.Username, user.Email, user.Password, user.Picture); err != nil {
		panic(err)
	}
}

func UpdateUser(user User) {
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

func DeleteUser(id string) {
	Connect()

	// Delete user picture from public
	if rows, err := DB.Query("SELECT picture FROM user WHERE id=?", id); err != nil {
		panic(err)
	} else {
		defer rows.Close()

		picture := ""
		for rows.Next() {
			if err := rows.Scan(&picture); err != nil {
				panic(err)
			}
		}
		if err := rows.Err(); err != nil {
			panic(err)
		}
		deleteImages(picture)
	}

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
			DeleteBlog(id)
		}(id)
	}
	wg.Wait()

	const sql = `DELETE FROM user WHERE id=?`

	if _, err := DB.Exec(sql, id); err != nil {
		panic(err)
	}
}

func SelectUser(id string) {
	Connect()
	const sql = ``

	rows, err := DB.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

}

func ProfileUser(id string) {
	Connect()
	const sql = ``

	rows, err := DB.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}

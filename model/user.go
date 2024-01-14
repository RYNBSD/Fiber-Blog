package model

func CreateUser(user User) {
	Connect()
	sql := `INSERT INTO user(username, email, password, picture) VALUES (?, ?, ?, ?)`

	if _, err := DB.Exec(sql, user.Username, user.Email, user.Password, user.Picture); err != nil {
		panic(err)
	}
}

func UpdateUser(user User) {
	Connect()
	sql := `UPDATE user SET username=?, email=?, password=?, picture=? WHERE id=?`

	if _, err := DB.Exec(sql, user.Username, user.Email, user.Password, user.Picture, user.Id); err != nil {
		panic(err)
	}
}

func DeleteUser(id string) {
	Connect()
	sql := `DELETE FROM user WHERE id=?`

	if _, err := DB.Exec(sql, id); err != nil {
		panic(err)
	}
}

func SelectUser(id string) {
	Connect()
	sql := ``

	rows, err := DB.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

}
func ProfileUser() {
	Connect()
	sql := ``

	rows, err := DB.Query(sql, id)
	if err != nil {
		panic(err)
	}
	defer rows.Close()
}

func LikeBlog(userId string, blogId string) {
	Connect()
	sql := `INSERT INTO "blogLikes" (likerId, blogId) VALUES (?, ?)`

	if _, err := DB.Exec(sql, userId, blogId); err != nil {
		panic(err)
	}
}

func CommentBlog(userId string, blogId string, comment string) {
	Connect()
	sql := `INSERT INTO "blogComments" (likerId, blogId, comment) VALUES (?, ?, ?)`

	if _, err := DB.Exec(sql, userId, blogId, comment); err != nil {
		panic(err)
	}
}

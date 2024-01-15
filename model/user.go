package model

func CreateUser(user User) {
	Connect()
	const sql = `INSERT INTO user(username, email, password, picture) VALUES (?, ?, ?, ?)`

	if _, err := DB.Exec(sql, user.Username, user.Email, user.Password, user.Picture); err != nil {
		panic(err)
	}
}

func UpdateUser(user User) {
	Connect()
	const sql = `UPDATE user SET username=?, email=?, password=?, picture=? WHERE id=?`

	if _, err := DB.Exec(sql, user.Username, user.Email, user.Password, user.Picture, user.Id); err != nil {
		panic(err)
	}
}

func DeleteUser(id string) {
	Connect()
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

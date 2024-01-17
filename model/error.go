package model

func (e *Error) CreateError() {
	Connect()
	sql := `INSERT INTO error(status, message) VALUES (?, ?)`

	if _, err := DB.Exec(sql, e.Status, e.Message); err != nil {
		panic(err)
	}
}

package schema

type Update struct {
	Username string `json:"username" validate:"required,min=1"`
	Email string `json:"email" validate:"required,min=1,email"`
}

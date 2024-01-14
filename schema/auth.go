package schema

// Request Body Schema For All Auth Endpoints

type SignUp struct {
	Username string `json:"username" validate:"required,min=1"`
	Email    string `json:"email" validate:"required,min=1,email"`
	Password string `json:"password" validate:"required,min=8"`
}

type SignIn struct {
	Email    string `json:"email" validate:"required,min=1,email"`
	Password string `json:"password" validate:"required,min=8"`
}

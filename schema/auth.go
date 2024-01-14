package schema

// Request Body Schema For All Auth Endpoints

type SignUp struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignIn struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

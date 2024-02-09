package schema

type CreateBlog struct {
	Title       string `json:"title" validate:"required,min=1"`
	Description string `json:"description" validate:"required,min=1"`
}

type CreateComment struct {
	Comment string `json:"comment" validate:"required,min=1"`
}

type UpdateBlog struct {
	Title       string `json:"title" validate:"required,min=1"`
	Description string `json:"description" validate:"required,min=1"`
}

type UpdateComment struct {
	Comment string `json:"comment" validate:"required,min=1"`
}

type DeleteBlog struct {
}

type DeleteComment struct {
}

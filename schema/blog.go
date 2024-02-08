package schema

type CreateBlog struct {
	Title string `json:"title" validate:"required,min=1"`
	Description string `json:"description" validate:"required,min=1"`
}
package models

//Books model
type Books struct {
	ID          string `form:"id" json:"id"`
	Title       string `form:"title" json:"title"`
	Description string `form:"description" json:"description"`
	AuthorID    string `form:"author_id" json:"author_id"`
}

//BookResponse model
type BookResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Books
}

package models

//Authors model
type Authors struct {
	ID   string `form:"id" json:"id"`
	Name string `form:"name" json:"name"`
	DOB  string `form:"dob" json:"dob"`
}

//AuthorResponse model
type AuthorResponse struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
	Data    []Authors
}

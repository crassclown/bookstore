package action

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/crassclown/bookstore/db"
	"github.com/crassclown/bookstore/models"
)

// ReturnAllAuthors function
func ReturnAllAuthors(w http.ResponseWriter, r *http.Request) {
	var authors models.Authors
	var arrAuthors []models.Authors
	var response models.AuthorResponse

	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("select id, name, dob from authors")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&authors.ID, &authors.Name, &authors.DOB); err != nil {
			log.Fatal(err.Error())
		} else {
			arrAuthors = append(arrAuthors, authors)
		}
	}

	if len(arrAuthors) > 0 {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 404
		response.Message = "Not Found"
	}
	response.Data = arrAuthors

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

// CreateAuthor function
func CreateAuthor(w http.ResponseWriter, r *http.Request) {
	var authors models.Authors
	var arrAuthors []models.Authors
	var response models.AuthorResponse

	db := db.Connect()
	defer db.Close()

	err := r.ParseMultipartForm(4096)
	if err != nil {
		panic(err)
	}

	name := r.FormValue("name")
	dob := r.FormValue("dob")

	_, err = db.Exec("insert into authors (name, dob) values (?,?)", name, dob)

	if err != nil {
		log.Print(err)
	}

	err = db.QueryRow("select LAST_INSERT_ID()").Scan(&authors.ID)
	if err != nil {
		response.Status = 400
		response.Message = "Bad Request"
	} else {
		response.Status = 200
		response.Message = "Success"
		response.Data = arrAuthors
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

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

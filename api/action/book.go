package action

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/crassclown/bookstore/db"
	"github.com/crassclown/bookstore/models"
)

//ReturnAllBooks function
func ReturnAllBooks(w http.ResponseWriter, r *http.Request) {
	var books models.Books
	var arrBooks []models.Books
	var response models.BookResponse

	db := db.Connect()
	defer db.Close()

	rows, err := db.Query("Select id, title, description, author_id from books")

	if err != nil {
		log.Print(err)
	}

	for rows.Next() {
		if err := rows.Scan(&books.ID, &books.Title, &books.Description, &books.AuthorID); err != nil {
			log.Fatal(err.Error())
		} else {
			arrBooks = append(arrBooks, books)
		}
	}

	if len(arrBooks) > 0 {
		response.Status = 200
		response.Message = "Success"
	} else {
		response.Status = 404
		response.Message = "Not Found"
	}
	response.Data = arrBooks

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

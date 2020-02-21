package main

import (
	"database/sql"
	"log"
)

//Connect DB function
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/bookstore")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

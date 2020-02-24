package db

import (
	"database/sql"
	"log"

	//mysql driver
	_ "github.com/go-sql-driver/mysql"
)

//Connect DB function
func Connect() *sql.DB {
	db, err := sql.Open("mysql", "root:@tcp(localhost:3306)/bookstore")

	if err != nil {
		log.Fatal(err)
	}

	return db
}

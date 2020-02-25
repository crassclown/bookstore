package router

import (
	"log"
	"net/http"

	"github.com/crassclown/bookstore/action"

	"github.com/gorilla/mux"
)

//Router function
func Router() {
	port := ":1234"
	router := NewRouter()
	log.Println("Magic happens on port", port)
	log.Fatal(http.ListenAndServe(port, router))
}

//NewRouter function
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/books", action.ReturnAllBooks).Methods("GET")
	router.HandleFunc("/authors", action.ReturnAllAuthors).Methods("GET")
	return router
}

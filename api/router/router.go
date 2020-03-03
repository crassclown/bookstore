package router

import (
	"log"
	"net/http"

	"github.com/crassclown/bookstore/action"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

//Router function
func Router() {
	port := ":1234"
	router := NewRouter()
	handler := cors.Default().Handler(router)
	log.Println("Magic happens on port", port)
	log.Fatal(http.ListenAndServe(port, handler))
}

//NewRouter function
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	// book routes
	router.HandleFunc("/books", action.ReturnAllBooks).Methods("GET")
	// router.HandleFunc("/books", action.CreateBook).Methods("POST")

	// author routes
	router.HandleFunc("/authors", action.ReturnAllAuthors).Methods("GET")
	router.HandleFunc("/authors", action.CreateAuthor).Methods("POST")
	return router
}

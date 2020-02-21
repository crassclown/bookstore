package router

import (
	"net/http"

	"github.com/crassclown/bookstore/action"

	"github.com/gorilla/mux"
)

//Router function
func Router() {
	router := NewRouter()
	http.ListenAndServe(":1234", router)
}

//NewRouter function
func NewRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)

	router.HandleFunc("/", action.ReturnAllBooks).Methods("GET")
	return router
}

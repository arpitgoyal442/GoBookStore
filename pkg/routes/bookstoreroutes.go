package routes

import (
	"github.com/gorilla/mux"
	"github.com/arpit/go-bookstore-mysql/pkg/controllers"
)

var MyBookStoreRoutes = func(router *mux.Router) {

	router.HandleFunc("/book/", controllers.GetBook).Methods("GET")

	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/{Id}", controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{Id}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{Id}", controllers.GetBookById).Methods("GET")

}

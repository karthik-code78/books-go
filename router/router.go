package router

import (
	"github.com/gorilla/mux"
	"library-management/handlers"
	"library-management/middlewares/logging"
)

func InitRouter() *mux.Router {
	router := mux.NewRouter()

	// middlewares
	router.Use(logging.Logger)

	// Books router
	booksRouter := router.PathPrefix("/books").Subrouter()

	// Books routes - CRUD
	booksRouter.HandleFunc("", handlers.GetAllBooks).Methods("GET")
	booksRouter.HandleFunc("/{id}", handlers.GetBook).Methods("GET")
	booksRouter.HandleFunc("", handlers.CreateBook).Methods("POST")
	booksRouter.HandleFunc("/{id}", handlers.UpdateBook).Methods("PUT")
	booksRouter.HandleFunc("/{id}", handlers.DeleteBook).Methods("DELETE")

	return router
}

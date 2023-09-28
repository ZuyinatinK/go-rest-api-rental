package routes

import (
	"go-rest-api-rental/controllers/propertycontroller"

	"github.com/gorilla/mux"
)

func PropertyRoutes(r *mux.Router) {
	router := r.PathPrefix("/properties").Subrouter()

	router.HandleFunc("", propertycontroller.Index).Methods("GET")
	router.HandleFunc("", propertycontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", propertycontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", propertycontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", propertycontroller.Delete).Methods("DELETE")
}

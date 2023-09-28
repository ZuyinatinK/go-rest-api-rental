package routes

import (
	"go-rest-api-rental/controllers/tenantcontroller"

	"github.com/gorilla/mux"
)

func TenantRoutes(r *mux.Router) {
	router := r.PathPrefix("/tenants").Subrouter()

	router.HandleFunc("", tenantcontroller.Index).Methods("GET")
	router.HandleFunc("", tenantcontroller.Create).Methods("POST")
	router.HandleFunc("/{id}/detail", tenantcontroller.Detail).Methods("GET")
	router.HandleFunc("/{id}/update", tenantcontroller.Update).Methods("PUT")
	router.HandleFunc("/{id}/delete", tenantcontroller.Delete).Methods("DELETE")
}

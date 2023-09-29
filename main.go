package main

import (
	"fmt"
	"go-rest-api-rental/config"
	"go-rest-api-rental/routes"

	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()
	routes.RouteIndex(r)

	log.Println("Server running")
	http.ListenAndServe(fmt.Sprintf(":%v", config.ENV.PORT), r)
}

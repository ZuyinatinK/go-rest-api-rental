package main

import (
	"fmt"
	"go-rest-api-rental/config"
	"go-rest-api-rental/routes"
	"os"

	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/gorilla/mux"
)

func main() {
	config.LoadEnv()
	config.LoadConfig()
	config.ConnectDB()

	r := mux.NewRouter()
	routes.RouteIndex(r)

	log.Println("Server running on port", os.Getenv("PORT"))
	http.ListenAndServe(fmt.Sprintf(":%v", os.Getenv("PORT")), r)
}

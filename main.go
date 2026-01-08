package main

import (
	"log"
	"net/http"
	"todo-api-golang/config"
	"todo-api-golang/routes"
)

func main() {
	config.ConnectDB()

	routes.SetupRoutes()

	log.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

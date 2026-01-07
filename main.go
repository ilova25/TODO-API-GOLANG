package main

import (
	"fmt"
	"net/http"
	"todo-api-golang/routes"
)

func main() {
	routes.SetupRoutes()

	fmt.Println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

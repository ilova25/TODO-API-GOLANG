package routes

import (
	"net/http"
	"todo-api-golang/handlers"
)

func SetupRoutes() {
	http.HandleFunc("/todos", handlers.GetTodos)
	http.HandleFunc("/todo", handlers.GetTodoByID)
	http.HandleFunc("/todo/create", handlers.CreateTodo)
	http.HandleFunc("/todo/update", handlers.UpdateTodo)
	http.HandleFunc("/todo/delete", handlers.DeleteTodo)
}

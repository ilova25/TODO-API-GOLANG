package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"todo-api-golang/models"
	"todo-api-golang/utils"
)

var todos []models.Todo
var idCounter = 1

// CREATE TODO
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	todo.ID = idCounter
	todo.Status = "pending"
	todo.CreatedAt = time.Now()

	idCounter++
	todos = append(todos, todo)

	utils.JSONResponse(w, http.StatusCreated, todo)
}

// GET LIST TODO
func GetTodos(w http.ResponseWriter, r *http.Request) {
	utils.JSONResponse(w, http.StatusOK, todos)
}

// GET DETAIL TODO
func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	for _, todo := range todos {
		if todo.ID == id {
			utils.JSONResponse(w, http.StatusOK, todo)
			return
		}
	}
	utils.JSONResponse(w, http.StatusNotFound, map[string]string{
		"message": "Todo not found",
	})
}

// UPDATE TODO
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var updatedTodo models.Todo
	json.NewDecoder(r.Body).Decode(&updatedTodo)

	for i, todo := range todos {
		if todo.ID == id {
			todos[i].Title = updatedTodo.Title
			todos[i].Description = updatedTodo.Description
			todos[i].Status = updatedTodo.Status

			utils.JSONResponse(w, http.StatusOK, todos[i])
			return
		}
	}
	utils.JSONResponse(w, http.StatusNotFound, map[string]string{
		"message": "Todo not found",
	})
}

// DELETE TODO
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	for i, todo := range todos {
		if todo.ID == id {
			todos = append(todos[:i], todos[i+1:]...)
			utils.JSONResponse(w, http.StatusOK, map[string]string{
				"message": "Todo deleted",
			})
			return
		}
	}
	utils.JSONResponse(w, http.StatusNotFound, map[string]string{
		"message": "Todo not found",
	})
}

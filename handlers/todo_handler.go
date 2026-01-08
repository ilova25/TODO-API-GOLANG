package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"todo-api-golang/config"
	"todo-api-golang/models"
	"todo-api-golang/utils"
)


// ================= CREATE TODO =================
func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	query := "INSERT INTO todos (title, description) VALUES (?, ?)"
	result, err := config.DB.Exec(query, todo.Title, todo.Description)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	id, _ := result.LastInsertId()
	todo.ID = int(id)
	todo.Status = "pending"

	utils.JSONResponse(w, http.StatusCreated, todo)
}


// ================= GET LIST TODO =================
func GetTodos(w http.ResponseWriter, r *http.Request) {
	rows, err := config.DB.Query(
		"SELECT id, title, description, status, created_at FROM todos",
	)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	defer rows.Close()

	var todos []models.Todo

	for rows.Next() {
		var todo models.Todo
		rows.Scan(
			&todo.ID,
			&todo.Title,
			&todo.Description,
			&todo.Status,
			&todo.CreatedAt,
		)
		todos = append(todos, todo)
	}

	utils.JSONResponse(w, http.StatusOK, todos)
}


// ================= GET DETAIL TODO =================
func GetTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var todo models.Todo
	err := config.DB.QueryRow(
		"SELECT id, title, description, status, created_at FROM todos WHERE id = ?",
		id,
	).Scan(
		&todo.ID,
		&todo.Title,
		&todo.Description,
		&todo.Status,
		&todo.CreatedAt,
	)

	if err != nil {
		utils.JSONResponse(w, http.StatusNotFound, map[string]string{
			"message": "Todo not found",
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, todo)
}


// ================= UPDATE TODO =================
func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)

	query := `
		UPDATE todos 
		SET title = ?, description = ?, status = ?
		WHERE id = ?
	`

	result, err := config.DB.Exec(
		query,
		todo.Title,
		todo.Description,
		todo.Status,
		id,
	)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		utils.JSONResponse(w, http.StatusNotFound, map[string]string{
			"message": "Todo not found",
		})
		return
	}

	todo.ID = id
	utils.JSONResponse(w, http.StatusOK, todo)
}


// ================= DELETE TODO =================
func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Query().Get("id")
	id, _ := strconv.Atoi(idStr)

	result, err := config.DB.Exec(
		"DELETE FROM todos WHERE id = ?",
		id,
	)
	if err != nil {
		utils.JSONResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	rows, _ := result.RowsAffected()
	if rows == 0 {
		utils.JSONResponse(w, http.StatusNotFound, map[string]string{
			"message": "Todo not found",
		})
		return
	}

	utils.JSONResponse(w, http.StatusOK, map[string]string{
		"message": "Todo deleted",
	})
}

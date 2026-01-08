# Todo List REST API (Golang)
Todo List REST API adalah aplikasi backend sederhana yang dibuat menggunakan **Golang**.  
Aplikasi ini menyediakan layanan **CRUD (Create, Read, Update, Delete)** untuk data Todo dan menggunakan **JSON** sebagai format response.
Project ini dibuat sebagai **Tes PKL**.

## 🚀 Fitur
- Create Todo
- Get List Todo
- Get Detail Todo
- Update Todo
- Delete Todo
- Penyimpanan data menggunakan MySQL
- Response JSON
- HTTP Status Code sesuai standar
- Struktur project rapi

## 🛠️ Teknologi
- Golang
- net/http
- MySQL
- go-sql-driver/mysql
- JSON

## Setup Database MySQL
CREATE DATABASE todo_db;
USE todo_db;

CREATE TABLE todos (
  id INT AUTO_INCREMENT PRIMARY KEY,
  title VARCHAR(255) NOT NULL,
  description TEXT,
  status ENUM('pending','done') DEFAULT 'pending',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


## ▶️ Cara Menjalankan Project

## 1. Clone repository:
git clone https://github.com/ilova25/TODO-API-GOLANG.git

## 2. Masuk ke folder project:
cd TODO-API-GOLANG

## 3. Jalankan server:
go run main.go

## 4. Server berjalan di:
http://localhost:8080

## 📂 Struktur Project
todo-api-golang
│── main.go
│── go.mod
├── config
│   └── db.go
├── handlers
│   └── todo_handler.go
├── models
│   └── todo.go
├── routes
│   └── routes.go
└── utils
    └── response.go


## 📌 Struktur Data Todo
{
  "id": 1,
  "title": "Belajar Golang",
  "description": "Membuat REST API Todo",
  "status": "pending",
  "created_at": "2026-01-07T12:00:00Z"
}

## 📮 API Endpoint
## 1️⃣ Create Todo
## POST /todo/create
## Request Body
{
  "title": "Belajar Golang",
  "description": "Membuat REST API Todo"
}
## Response (201 Created)
{
  "id": 1,
  "title": "Belajar Golang",
  "description": "Membuat REST API Todo",
  "status": "pending",
  "created_at": "2026-01-07T12:00:00Z"
}

## 2️⃣ Get List Todo
## GET /todos
## Response (200 OK)
[
  {
    "id": 1,
    "title": "Belajar Golang",
    "description": "Membuat REST API Todo",
    "status": "pending",
    "created_at": "2026-01-07T12:00:00Z"
  }
]

## 3️⃣ Get Detail Todo
## GET /todo?id=1
## Response (200 OK)
{
  "id": 1,
  "title": "Belajar Golang",
  "description": "Membuat REST API Todo",
  "status": "pending",
  "created_at": "2026-01-07T12:00:00Z"
}
##  Response (404 Not Found)
{
  "message": "Todo not found"
}

## 4️⃣ Update Todo
## PUT /todo/update?id=1
## Request Body
{
  "title": "Belajar Golang REST",
  "description": "Update data todo",
  "status": "done"
}
## Response (200 OK)
{
  "id": 1,
  "title": "Belajar Golang REST",
  "description": "Update data todo",
  "status": "done",
  "created_at": "2026-01-07T12:00:00Z"
}

## 5️⃣ Delete Todo
## DELETE /todo/delete?id=1
## Response (200 OK)
{
  "message": "Todo deleted"
}

## 🧪 Testing
API dapat diuji menggunakan:
Postman
Browser (GET request)
cURL

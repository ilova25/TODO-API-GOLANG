package models

import "time"

type Todo struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Status      string    `json:"status"` // done / pending
	CreatedAt   time.Time `json:"created_at"`
}

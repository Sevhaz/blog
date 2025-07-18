package models

import "time"

type Blog struct {
	ID        string    `json:"id"`
	UserID    string    `json:"user_id"` //we need to know who created the todo- this will identify the id belongs to the person
	Title     string    `json:"title"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	DeletedAt time.Time `json:"deleted_at"`
}


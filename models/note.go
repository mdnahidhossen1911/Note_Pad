package models

import "time"

type Note struct {
	ID        string    `json:"id" gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UID       string    `json:"uid"     binding:"required"`
	Title     string    `json:"title"`
	Body      string    `josn:"body"`
	CreatedAt time.Time `json:"created_at" gorm:"autoCreateTime"`
	UpdatedAt time.Time `json:"updated_at" gorm:"autoUpdateTime"`
}

type NoteRequest struct {
	Title string `json:"title"  binding:"required"`
	Body  string `josn:"body"  binding:"required"`
}

type NoteUpdateRequest struct {
	ID    string `json:"id" binding:"required"`
	Title string `json:"title"`
	Body  string `josn:"body"`
}

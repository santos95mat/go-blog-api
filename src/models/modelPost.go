package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Post struct {
	ID     uuid.UUID `gorm:"primaryKey"`
	Title  string    `json:"title" gorm:"not null"`
	Body   string    `json:"body" gorm:"not null"`
	UserID uuid.UUID `json:"user_id" db:"user_id"`
	gorm.Model
}

type PostDto struct {
	Title string `json:"title"`
	Body  string `json:"body"`
}

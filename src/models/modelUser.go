package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID       uuid.UUID `gorm:"primaryKey"`
	Name     string    `json:"name" gorm:"not null"`
	LastName string    `json:"last_name" gorm:"not null" db:"last_name"`
	Email    string    `json:"email" gorm:"not null;unique_index"`
	Password string    `json:"-" gorm:"not null"`
	Nickname string    `json:"nickname" gorm:"not null;unique_index;"`
	Posts    []Post    `json:"posts" gorm:"foreignKey:UserID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	gorm.Model
}

type UserDto struct {
	Name     string `json:"name"`
	LastName string `json:"last_name" db:"last_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Nickname string `json:"nickname"`
}

type UserAuthDto struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserToken struct {
	User  User
	Token string
}

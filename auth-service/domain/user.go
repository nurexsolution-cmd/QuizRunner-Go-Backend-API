package domain

import (
	"time"
)

type User struct {
	ID           uint      `json:"id" gorm:"primaryKey"`
	Username     string    `json:"username" gorm:"uniqueIndex;not null"`
	Email        string    `json:"email" gorm:"uniqueIndex;not null"`
	PasswordHash string    `json:"-" gorm:"not null"` // Don't return password hash in JSON
	CreatedAt    time.Time `json:"created_at"`
}

type AuthUseCase interface {
	Register(username, email, password string) (*User, error)
	Login(email, password string) (string, error)
}

type AuthRepository interface {
	CreateUser(user *User) error
	GetUserByEmail(email string) (*User, error)
}

package domain

import (
	"time"
)

type User struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	Username  string    `json:"username"`
	Email     string    `json:"email"`
	CreatedAt time.Time `json:"created_at"`
}

type UserProgress struct {
	TotalQuizzesTaken int `json:"total_quizzes_taken"`
	TotalScore        int `json:"total_score"`
}

type UserRanking struct {
	UserID   uint   `json:"user_id"`
	Username string `json:"username"`
	Score    int    `json:"score"`
}

type UserUseCase interface {
	GetProfile(userID uint) (*User, error)
	GetProgress(userID uint) (*UserProgress, error)
	GetRankings() ([]UserRanking, error)
}

type UserRepository interface {
	GetUserByID(id uint) (*User, error)
	GetUserProgress(id uint) (*UserProgress, error)
	GetRankings() ([]UserRanking, error)
}

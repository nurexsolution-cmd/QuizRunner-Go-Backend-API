package repository

import (
	"errors"

	"github.com/pubudulakmal/quiz-backend/auth-service/domain"
	"gorm.io/gorm"
)

type pgAuthRepository struct {
	db *gorm.DB
}

func NewPostgresAuthRepository(db *gorm.DB) domain.AuthRepository {
	return &pgAuthRepository{db: db}
}

func (r *pgAuthRepository) CreateUser(user *domain.User) error {
	return r.db.Create(user).Error
}

func (r *pgAuthRepository) GetUserByEmail(email string) (*domain.User, error) {
	var user domain.User
	err := r.db.Where("email = ?", email).First(&user).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil // Return nil, nil when not found to easily check
	}
	if err != nil {
		return nil, err
	}
	return &user, nil
}

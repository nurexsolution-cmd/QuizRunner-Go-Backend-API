package usecase

import (
	"github.com/pubudulakmal/quiz-backend/user-service/domain"
)

type userUseCase struct {
	repo domain.UserRepository
}

func NewUserUseCase(r domain.UserRepository) domain.UserUseCase {
	return &userUseCase{repo: r}
}

func (u *userUseCase) GetProfile(userID uint) (*domain.User, error) {
	return u.repo.GetUserByID(userID)
}

func (u *userUseCase) GetProgress(userID uint) (*domain.UserProgress, error) {
	return u.repo.GetUserProgress(userID)
}

func (u *userUseCase) GetRankings() ([]domain.UserRanking, error) {
	return u.repo.GetRankings()
}

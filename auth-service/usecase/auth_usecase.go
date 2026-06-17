package usecase

import (
	"errors"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/pubudulakmal/quiz-backend/auth-service/domain"
	"golang.org/x/crypto/bcrypt"
)

type authUseCase struct {
	repo domain.AuthRepository
}

func NewAuthUseCase(r domain.AuthRepository) domain.AuthUseCase {
	return &authUseCase{repo: r}
}

func (u *authUseCase) Register(username, email, password string) (*domain.User, error) {
	existingUser, err := u.repo.GetUserByEmail(email)
	if err != nil {
		return nil, err
	}
	if existingUser != nil {
		return nil, errors.New("email already in use")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	user := &domain.User{
		Username:     username,
		Email:        email,
		PasswordHash: string(hashedPassword),
		CreatedAt:    time.Now(),
	}

	err = u.repo.CreateUser(user)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *authUseCase) Login(email, password string) (string, error) {
	user, err := u.repo.GetUserByEmail(email)
	if err != nil {
		return "", err
	}
	if user == nil {
		return "", errors.New("invalid email or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	if err != nil {
		return "", errors.New("invalid email or password")
	}

	// Generate JWT
	secret := os.Getenv("JWT_SECRET")
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

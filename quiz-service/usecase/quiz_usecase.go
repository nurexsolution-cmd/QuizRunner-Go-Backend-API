package usecase

import (
	"github.com/pubudulakmal/quiz-backend/quiz-service/domain"
)

type quizUseCase struct {
	repo domain.QuizRepository
}

func NewQuizUseCase(r domain.QuizRepository) domain.QuizUseCase {
	return &quizUseCase{repo: r}
}

func (u *quizUseCase) GetSubjects() ([]domain.Subject, error) {
	return u.repo.GetSubjects()
}

func (u *quizUseCase) GetQuizzesBySubject(subjectID uint) ([]domain.Quiz, error) {
	return u.repo.GetQuizzesBySubject(subjectID)
}

func (u *quizUseCase) GetQuiz(quizID uint) (*domain.Quiz, error) {
	return u.repo.GetQuiz(quizID)
}

func (u *quizUseCase) CreateSubject(subject *domain.Subject) error {
	return u.repo.CreateSubject(subject)
}

func (u *quizUseCase) CreateQuiz(quiz *domain.Quiz) error {
	return u.repo.CreateQuiz(quiz)
}

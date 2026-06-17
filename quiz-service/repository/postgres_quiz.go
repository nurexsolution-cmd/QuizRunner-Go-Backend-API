package repository

import (
	"github.com/pubudulakmal/quiz-backend/quiz-service/domain"
	"gorm.io/gorm"
)

type pgQuizRepository struct {
	db *gorm.DB
}

func NewPostgresQuizRepository(db *gorm.DB) domain.QuizRepository {
	return &pgQuizRepository{db: db}
}

func (r *pgQuizRepository) GetSubjects() ([]domain.Subject, error) {
	var subjects []domain.Subject
	err := r.db.Find(&subjects).Error
	return subjects, err
}

func (r *pgQuizRepository) GetQuizzesBySubject(subjectID uint) ([]domain.Quiz, error) {
	var quizzes []domain.Quiz
	err := r.db.Where("subject_id = ?", subjectID).Find(&quizzes).Error
	return quizzes, err
}

func (r *pgQuizRepository) GetQuiz(quizID uint) (*domain.Quiz, error) {
	var quiz domain.Quiz
	err := r.db.Preload("Questions.Answers").First(&quiz, quizID).Error
	if err != nil {
		return nil, err
	}
	return &quiz, nil
}

func (r *pgQuizRepository) CreateSubject(subject *domain.Subject) error {
	return r.db.Create(subject).Error
}

func (r *pgQuizRepository) CreateQuiz(quiz *domain.Quiz) error {
	return r.db.Create(quiz).Error
}

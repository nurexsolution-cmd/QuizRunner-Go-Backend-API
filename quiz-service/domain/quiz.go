package domain

import (
	"time"
)

type Subject struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
}

type Quiz struct {
	ID          uint       `json:"id" gorm:"primaryKey"`
	SubjectID   uint       `json:"subject_id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	CreatedAt   time.Time  `json:"created_at"`
	Questions   []Question `json:"questions" gorm:"foreignKey:QuizID"`
}

type Question struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	QuizID    uint      `json:"quiz_id"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	Answers   []Answer  `json:"answers" gorm:"foreignKey:QuestionID"`
}

type Answer struct {
	ID         uint   `json:"id" gorm:"primaryKey"`
	QuestionID uint   `json:"question_id"`
	Text       string `json:"text"`
	IsCorrect  bool   `json:"is_correct"` // Note: In a real app we might omit this in JSON for students, but keeping it for practice
}

type QuizUseCase interface {
	GetSubjects() ([]Subject, error)
	GetQuizzesBySubject(subjectID uint) ([]Quiz, error)
	GetQuiz(quizID uint) (*Quiz, error)
	CreateSubject(subject *Subject) error
	CreateQuiz(quiz *Quiz) error
}

type QuizRepository interface {
	GetSubjects() ([]Subject, error)
	GetQuizzesBySubject(subjectID uint) ([]Quiz, error)
	GetQuiz(quizID uint) (*Quiz, error)
	CreateSubject(subject *Subject) error
	CreateQuiz(quiz *Quiz) error
}

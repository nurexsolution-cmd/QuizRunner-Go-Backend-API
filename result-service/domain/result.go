package domain

import "time"

type Result struct {
	ID             uint      `json:"id" gorm:"primaryKey"`
	UserID         uint      `json:"user_id"`
	QuizID         uint      `json:"quiz_id"`
	Score          int       `json:"score"`
	TotalQuestions int       `json:"total_questions"`
	CreatedAt      time.Time `json:"created_at"`
}

type UserAnswer struct {
	ID         uint `json:"id" gorm:"primaryKey"`
	ResultID   uint `json:"result_id"`
	QuestionID uint `json:"question_id"`
	AnswerID   uint `json:"answer_id"`
	IsCorrect  bool `json:"is_correct"`
}

type QuizSubmission struct {
	QuizID  uint           `json:"quiz_id" binding:"required"`
	Answers []SubmitAnswer `json:"answers" binding:"required"`
}

type SubmitAnswer struct {
	QuestionID uint `json:"question_id"`
	AnswerID   uint `json:"answer_id"`
}

type ResultUseCase interface {
	SubmitQuiz(userID uint, submission *QuizSubmission) (*Result, error)
	GetWrongAnswers(resultID uint) ([]UserAnswer, error)
}

type ResultRepository interface {
	SaveResult(result *Result, userAnswers []UserAnswer) error
	CheckAnswerIsCorrect(questionID, answerID uint) (bool, error)
	CountQuestionsInQuiz(quizID uint) (int, error)
	GetWrongAnswers(resultID uint) ([]UserAnswer, error)
	GetResultByIDAndUser(resultID, userID uint) (*Result, error)
}

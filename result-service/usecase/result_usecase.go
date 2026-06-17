package usecase

import (
	"time"

	"github.com/pubudulakmal/quiz-backend/result-service/domain"
)

type resultUseCase struct {
	repo domain.ResultRepository
}

func NewResultUseCase(r domain.ResultRepository) domain.ResultUseCase {
	return &resultUseCase{repo: r}
}

func (u *resultUseCase) SubmitQuiz(userID uint, submission *domain.QuizSubmission) (*domain.Result, error) {
	score := 0
	var userAnswers []domain.UserAnswer

	for _, ans := range submission.Answers {
		isCorrect, err := u.repo.CheckAnswerIsCorrect(ans.QuestionID, ans.AnswerID)
		if err != nil {
			continue // Handle or skip missing answer records
		}
		if isCorrect {
			score++
		}
		userAnswers = append(userAnswers, domain.UserAnswer{
			QuestionID: ans.QuestionID,
			AnswerID:   ans.AnswerID,
			IsCorrect:  isCorrect,
		})
	}

	totalQuestions, _ := u.repo.CountQuestionsInQuiz(submission.QuizID)

	result := &domain.Result{
		UserID:         userID,
		QuizID:         submission.QuizID,
		Score:          score,
		TotalQuestions: totalQuestions,
		CreatedAt:      time.Now(),
	}

	err := u.repo.SaveResult(result, userAnswers)
	if err != nil {
		return nil, err
	}

	return result, nil
}

func (u *resultUseCase) GetWrongAnswers(resultID uint) ([]domain.UserAnswer, error) {
	return u.repo.GetWrongAnswers(resultID)
}

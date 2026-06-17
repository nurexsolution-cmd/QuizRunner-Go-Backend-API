package repository

import (
	"github.com/pubudulakmal/quiz-backend/result-service/domain"
	"gorm.io/gorm"
)

type pgResultRepository struct {
	db *gorm.DB
}

func NewPostgresResultRepository(db *gorm.DB) domain.ResultRepository {
	return &pgResultRepository{db: db}
}

func (r *pgResultRepository) SaveResult(result *domain.Result, userAnswers []domain.UserAnswer) error {
	return r.db.Transaction(func(tx *gorm.DB) error {
		if err := tx.Create(result).Error; err != nil {
			return err
		}
		for i := range userAnswers {
			userAnswers[i].ResultID = result.ID
		}
		if len(userAnswers) > 0 {
			if err := tx.Create(&userAnswers).Error; err != nil {
				return err
			}
		}
		return nil
	})
}

func (r *pgResultRepository) CheckAnswerIsCorrect(questionID, answerID uint) (bool, error) {
	var isCorrect bool
	err := r.db.Table("answers").Select("is_correct").Where("id = ? AND question_id = ?", answerID, questionID).Row().Scan(&isCorrect)
	return isCorrect, err
}

func (r *pgResultRepository) CountQuestionsInQuiz(quizID uint) (int, error) {
	var count int64
	err := r.db.Table("questions").Where("quiz_id = ?", quizID).Count(&count).Error
	return int(count), err
}

func (r *pgResultRepository) GetWrongAnswers(resultID uint) ([]domain.UserAnswer, error) {
	var answers []domain.UserAnswer
	err := r.db.Where("result_id = ? AND is_correct = false", resultID).Find(&answers).Error
	return answers, err
}

func (r *pgResultRepository) GetResultByIDAndUser(resultID, userID uint) (*domain.Result, error) {
	var res domain.Result
	err := r.db.Where("id = ? AND user_id = ?", resultID, userID).First(&res).Error
	if err != nil {
		return nil, err
	}
	return &res, nil
}

package repository

import (
	"database/sql"
	"github.com/pubudulakmal/quiz-backend/user-service/domain"
	"gorm.io/gorm"
)

type pgUserRepository struct {
	db *gorm.DB
}

func NewPostgresUserRepository(db *gorm.DB) domain.UserRepository {
	return &pgUserRepository{db: db}
}

func (r *pgUserRepository) GetUserByID(id uint) (*domain.User, error) {
	var user domain.User
	err := r.db.First(&user, id).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *pgUserRepository) GetUserProgress(id uint) (*domain.UserProgress, error) {
	var progress domain.UserProgress
	
	row := r.db.Table("results").Where("user_id = ?", id).Select("count(id) as total_quizzes_taken, sum(score) as total_score").Row()
	
	var totalScore sql.NullInt64
	var totalQuizzes sql.NullInt64
	
	err := row.Scan(&totalQuizzes, &totalScore)
	if err != nil {
		return &domain.UserProgress{TotalQuizzesTaken: 0, TotalScore: 0}, nil
	}
	
	progress.TotalQuizzesTaken = int(totalQuizzes.Int64)
	progress.TotalScore = int(totalScore.Int64)
	
	return &progress, nil
}

func (r *pgUserRepository) GetRankings() ([]domain.UserRanking, error) {
	var rankings []domain.UserRanking
	err := r.db.Table("users").
		Select("users.id as user_id, users.username, coalesce(sum(results.score), 0) as score").
		Joins("left join results on users.id = results.user_id").
		Group("users.id, users.username").
		Order("score desc").
		Limit(10).
		Scan(&rankings).Error
		
	if err != nil {
		return nil, err
	}
	return rankings, nil
}

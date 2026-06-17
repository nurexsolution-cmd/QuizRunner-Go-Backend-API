package repository

import (
	"github.com/pubudulakmal/quiz-backend/notification-service/domain"
	"gorm.io/gorm"
)

type pgNotificationRepository struct {
	db *gorm.DB
}

func NewPostgresNotificationRepository(db *gorm.DB) domain.NotificationRepository {
	return &pgNotificationRepository{db: db}
}

func (r *pgNotificationRepository) GetByUserID(userID uint) ([]domain.Notification, error) {
	var notifications []domain.Notification
	err := r.db.Where("user_id = ?", userID).Order("created_at desc").Find(&notifications).Error
	return notifications, err
}

func (r *pgNotificationRepository) MarkAsRead(notificationID, userID uint) error {
	return r.db.Model(&domain.Notification{}).
		Where("id = ? AND user_id = ?", notificationID, userID).
		Update("is_read", true).Error
}

func (r *pgNotificationRepository) Create(notification *domain.Notification) error {
	return r.db.Create(notification).Error
}

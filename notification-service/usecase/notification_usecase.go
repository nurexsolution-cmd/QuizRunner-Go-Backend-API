package usecase

import (
	"time"

	"github.com/pubudulakmal/quiz-backend/notification-service/domain"
)

type notificationUseCase struct {
	repo domain.NotificationRepository
}

func NewNotificationUseCase(r domain.NotificationRepository) domain.NotificationUseCase {
	return &notificationUseCase{repo: r}
}

func (u *notificationUseCase) GetUserNotifications(userID uint) ([]domain.Notification, error) {
	return u.repo.GetByUserID(userID)
}

func (u *notificationUseCase) MarkAsRead(notificationID, userID uint) error {
	return u.repo.MarkAsRead(notificationID, userID)
}

func (u *notificationUseCase) CreateNotification(userID uint, message string) error {
	notif := &domain.Notification{
		UserID:    userID,
		Message:   message,
		IsRead:    false,
		CreatedAt: time.Now(),
	}
	return u.repo.Create(notif)
}

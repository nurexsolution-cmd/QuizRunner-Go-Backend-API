package domain

import "time"

type Notification struct {
	ID        uint      `json:"id" gorm:"primaryKey"`
	UserID    uint      `json:"user_id"`
	Message   string    `json:"message"`
	IsRead    bool      `json:"is_read" gorm:"default:false"`
	CreatedAt time.Time `json:"created_at"`
}

type NotificationUseCase interface {
	GetUserNotifications(userID uint) ([]Notification, error)
	MarkAsRead(notificationID, userID uint) error
	CreateNotification(userID uint, message string) error
}

type NotificationRepository interface {
	GetByUserID(userID uint) ([]Notification, error)
	MarkAsRead(notificationID, userID uint) error
	Create(notification *Notification) error
}

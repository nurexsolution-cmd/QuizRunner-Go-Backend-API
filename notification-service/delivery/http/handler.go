package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pubudulakmal/quiz-backend/notification-service/domain"
	"github.com/pubudulakmal/quiz-backend/pkg/middleware"
)

type NotificationHandler struct {
	notificationUseCase domain.NotificationUseCase
}

func NewNotificationHandler(r *gin.Engine, us domain.NotificationUseCase) {
	handler := &NotificationHandler{
		notificationUseCase: us,
	}

	protected := r.Group("/notifications")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/", handler.GetNotifications)
		protected.PATCH("/:id/read", handler.MarkAsRead)
		// Internal endpoint to create notification
		protected.POST("/", handler.CreateNotification)
	}
}

func (h *NotificationHandler) GetNotifications(c *gin.Context) {
	userIDFloat, _ := c.Get("user_id")
	userID := uint(userIDFloat.(float64))

	notifications, err := h.notificationUseCase.GetUserNotifications(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, notifications)
}

func (h *NotificationHandler) MarkAsRead(c *gin.Context) {
	userIDFloat, _ := c.Get("user_id")
	userID := uint(userIDFloat.(float64))

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid notification ID"})
		return
	}

	if err := h.notificationUseCase.MarkAsRead(uint(id), userID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "marked as read"})
}

type CreateReq struct {
	UserID  uint   `json:"user_id" binding:"required"`
	Message string `json:"message" binding:"required"`
}

func (h *NotificationHandler) CreateNotification(c *gin.Context) {
	var req CreateReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.notificationUseCase.CreateNotification(req.UserID, req.Message); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "created"})
}

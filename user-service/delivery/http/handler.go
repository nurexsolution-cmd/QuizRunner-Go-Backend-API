package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pubudulakmal/quiz-backend/pkg/middleware"
	"github.com/pubudulakmal/quiz-backend/user-service/domain"
)

type UserHandler struct {
	userUseCase domain.UserUseCase
}

func NewUserHandler(r *gin.Engine, us domain.UserUseCase) {
	handler := &UserHandler{
		userUseCase: us,
	}
	
	protected := r.Group("/users")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/profile", handler.GetProfile)
		protected.GET("/progress", handler.GetProgress)
		protected.GET("/rankings", handler.GetRankings)
	}
}

func (h *UserHandler) GetProfile(c *gin.Context) {
	userIDFloat, _ := c.Get("user_id")
	userID := uint(userIDFloat.(float64))

	user, err := h.userUseCase.GetProfile(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (h *UserHandler) GetProgress(c *gin.Context) {
	userIDFloat, _ := c.Get("user_id")
	userID := uint(userIDFloat.(float64))

	progress, err := h.userUseCase.GetProgress(userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, progress)
}

func (h *UserHandler) GetRankings(c *gin.Context) {
	rankings, err := h.userUseCase.GetRankings()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, rankings)
}

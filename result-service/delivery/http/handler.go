package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pubudulakmal/quiz-backend/pkg/middleware"
	"github.com/pubudulakmal/quiz-backend/result-service/domain"
)

type ResultHandler struct {
	resultUseCase domain.ResultUseCase
}

func NewResultHandler(r *gin.Engine, us domain.ResultUseCase) {
	handler := &ResultHandler{
		resultUseCase: us,
	}

	protected := r.Group("/results")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.POST("/submit", handler.SubmitQuiz)
		protected.GET("/:id/wrong-answers", handler.GetWrongAnswers)
	}
}

func (h *ResultHandler) SubmitQuiz(c *gin.Context) {
	userIDFloat, _ := c.Get("user_id")
	userID := uint(userIDFloat.(float64))

	var sub domain.QuizSubmission
	if err := c.ShouldBindJSON(&sub); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.resultUseCase.SubmitQuiz(userID, &sub)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, result)
}

func (h *ResultHandler) GetWrongAnswers(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid result ID"})
		return
	}

	wrongAnswers, err := h.resultUseCase.GetWrongAnswers(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, wrongAnswers)
}

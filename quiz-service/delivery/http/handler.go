package http

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/pubudulakmal/quiz-backend/pkg/middleware"
	"github.com/pubudulakmal/quiz-backend/quiz-service/domain"
)

type QuizHandler struct {
	quizUseCase domain.QuizUseCase
}

func NewQuizHandler(r *gin.Engine, us domain.QuizUseCase) {
	handler := &QuizHandler{
		quizUseCase: us,
	}

	protected := r.Group("/quizzes")
	protected.Use(middleware.AuthMiddleware())
	{
		protected.GET("/subjects", handler.GetSubjects)
		protected.POST("/subjects", handler.CreateSubject)
		
		protected.GET("/subjects/:id", handler.GetQuizzesBySubject)
		protected.POST("/", handler.CreateQuiz)
		
		protected.GET("/:id", handler.GetQuiz)
	}
}

func (h *QuizHandler) GetSubjects(c *gin.Context) {
	subjects, err := h.quizUseCase.GetSubjects()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, subjects)
}

func (h *QuizHandler) CreateSubject(c *gin.Context) {
	var subject domain.Subject
	if err := c.ShouldBindJSON(&subject); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.quizUseCase.CreateSubject(&subject); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, subject)
}

func (h *QuizHandler) GetQuizzesBySubject(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid subject ID"})
		return
	}
	quizzes, err := h.quizUseCase.GetQuizzesBySubject(uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, quizzes)
}

func (h *QuizHandler) CreateQuiz(c *gin.Context) {
	var quiz domain.Quiz
	if err := c.ShouldBindJSON(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := h.quizUseCase.CreateQuiz(&quiz); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, quiz)
}

func (h *QuizHandler) GetQuiz(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid quiz ID"})
		return
	}
	quiz, err := h.quizUseCase.GetQuiz(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Quiz not found"})
		return
	}
	c.JSON(http.StatusOK, quiz)
}

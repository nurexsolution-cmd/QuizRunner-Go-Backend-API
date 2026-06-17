package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pubudulakmal/quiz-backend/pkg/db"
	"github.com/pubudulakmal/quiz-backend/quiz-service/delivery/http"
	"github.com/pubudulakmal/quiz-backend/quiz-service/domain"
	"github.com/pubudulakmal/quiz-backend/quiz-service/repository"
	"github.com/pubudulakmal/quiz-backend/quiz-service/usecase"
)

func main() {
	database := db.InitDB()

	database.AutoMigrate(&domain.Subject{}, &domain.Quiz{}, &domain.Question{}, &domain.Answer{})

	r := gin.Default()

	quizRepo := repository.NewPostgresQuizRepository(database)
	quizUC := usecase.NewQuizUseCase(quizRepo)
	http.NewQuizHandler(r, quizUC)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Quiz Service starting on port %s", port)
	r.Run(":" + port)
}

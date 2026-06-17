package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pubudulakmal/quiz-backend/pkg/db"
	"github.com/pubudulakmal/quiz-backend/result-service/delivery/http"
	"github.com/pubudulakmal/quiz-backend/result-service/domain"
	"github.com/pubudulakmal/quiz-backend/result-service/repository"
	"github.com/pubudulakmal/quiz-backend/result-service/usecase"
)

func main() {
	database := db.InitDB()

	database.AutoMigrate(&domain.Result{}, &domain.UserAnswer{})

	r := gin.Default()

	resultRepo := repository.NewPostgresResultRepository(database)
	resultUC := usecase.NewResultUseCase(resultRepo)
	http.NewResultHandler(r, resultUC)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Result Service starting on port %s", port)
	r.Run(":" + port)
}

package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pubudulakmal/quiz-backend/pkg/db"
	"github.com/pubudulakmal/quiz-backend/user-service/delivery/http"
	"github.com/pubudulakmal/quiz-backend/user-service/repository"
	"github.com/pubudulakmal/quiz-backend/user-service/usecase"
)

func main() {
	database := db.InitDB()

	r := gin.Default()

	userRepo := repository.NewPostgresUserRepository(database)
	userUC := usecase.NewUserUseCase(userRepo)
	http.NewUserHandler(r, userUC)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("User Service starting on port %s", port)
	r.Run(":" + port)
}

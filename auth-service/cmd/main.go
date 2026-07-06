package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pubudulakmal/quiz-backend/auth-service/db"
	"github.com/pubudulakmal/quiz-backend/auth-service/delivery/http"
	"github.com/pubudulakmal/quiz-backend/auth-service/domain"
	"github.com/pubudulakmal/quiz-backend/auth-service/repository"
	"github.com/pubudulakmal/quiz-backend/auth-service/usecase"
)

func main() {
	database := db.InitDB()

	// Auto migrate
	database.AutoMigrate(&domain.User{})

	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"service": "auth-service",
			"status":  "ok",
		})
	})
	r.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "ok"})
	})

	authRepo := repository.NewPostgresAuthRepository(database)
	authUC := usecase.NewAuthUseCase(authRepo)
	http.NewAuthHandler(r, authUC)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	
	log.Printf("Auth Service starting on port %s", port)
	r.Run(":" + port)
}

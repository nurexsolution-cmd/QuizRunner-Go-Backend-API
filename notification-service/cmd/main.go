package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/pubudulakmal/quiz-backend/notification-service/delivery/http"
	"github.com/pubudulakmal/quiz-backend/notification-service/domain"
	"github.com/pubudulakmal/quiz-backend/notification-service/repository"
	"github.com/pubudulakmal/quiz-backend/notification-service/usecase"
	"github.com/pubudulakmal/quiz-backend/pkg/db"
)

func main() {
	database := db.InitDB()
	database.AutoMigrate(&domain.Notification{})

	r := gin.Default()

	notifRepo := repository.NewPostgresNotificationRepository(database)
	notifUC := usecase.NewNotificationUseCase(notifRepo)
	http.NewNotificationHandler(r, notifUC)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("Notification Service starting on port %s", port)
	r.Run(":" + port)
}

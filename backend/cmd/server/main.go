package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ozoli99/Praesto/pkg/user"
	"github.com/ozoli99/Praesto/pkg/config"
	"github.com/ozoli99/Praesto/pkg/integration"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	configuration := config.Load()

	database, err := gorm.Open(postgres.Open(configuration.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	database.AutoMigrate(&user.User{})
	userRepository := user.NewGormRepository(database)
	userService := user.NewService(userRepository)

	router := gin.Default()
	router.Use(integration.AuthMiddleware())

	router.GET("/profile", func(context *gin.Context) {
		claimsInterface, exists := context.Get(integration.ClaimsKey)
		if !exists {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "no claims found"})
			return
		}
		claims, ok := claimsInterface.(map[string]interface{})
		if !ok {
			context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid claims"})
			return
		}

		user, err := userService.SyncUserFromClaims(claims)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, user)
	})

	if err := router.Run(":" + configuration.Port); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
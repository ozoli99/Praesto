package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ozoli99/Praesto/pkg/appointment"
	"github.com/ozoli99/Praesto/pkg/config"
	"github.com/ozoli99/Praesto/pkg/middleware"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	configuration := config.Load()

	database, err := gorm.Open(postgres.Open(configuration.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	appointmentRepository := appointment.NewGormRepository(database)
	appointmentService := appointment.NewService(appointmentRepository)

	router := gin.New()
	router.Use(middleware.LoggingMiddleware())

	router.GET("/test-appointment", func(context *gin.Context) {
		sampleAppointment := &appointment.Appointment{
			UserID:     1,
			ProviderID: 2,
			Category:   "gym",
			TimeSlot:   appointment.TimeNow(),
			Status:     "",
		}
		err := appointmentService.BookAppointment(sampleAppointment)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, gin.H{"message": "Appointment booked", "appointment": sampleAppointment})
	})

	if err := router.Run(":" + configuration.Port); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}
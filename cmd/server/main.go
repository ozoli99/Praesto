package main

import (
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/ozoli99/Praesto/appointment"
	"github.com/ozoli99/Praesto/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	configuration := config.Load()

	database, err := gorm.Open(postgres.Open(configuration.DatabaseURL), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	database.AutoMigrate(&appointment.Appointment{})

	appointmentRepository := appointment.NewGormRepository(database)
	appointmentService := appointment.NewService(appointmentRepository)

	router := gin.Default()

	router.POST("/appointments", func(context *gin.Context) {
		var request struct {
			ProviderID uint   `json:"provider_id"`
			CustomerID uint   `json:"customer_id"`
			StartTime  string `json:"start_time"`
			EndTime    string `json:"end_time"`
		}
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		startTime, err := time.Parse(time.RFC3339, request.StartTime)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid start time"})
			return
		}
		endTime, err := time.Parse(time.RFC3339, request.EndTime)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid end time"})
			return
		}
		appointment, err := appointmentService.BookAppointment(request.ProviderID, request.CustomerID, startTime, endTime)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusCreated, appointment)
	})

	router.PUT("/appointments/:id/reschedule", func(context *gin.Context) {
		idStr := context.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid appointment id"})
			return
		}
		var request struct {
			NewStartTime string `json:"new_start_time"`
			NewEndTime   string `json:"new_end_time"`
		}
		if err := context.ShouldBindJSON(&request); err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		newStartTime, err := time.Parse(time.RFC3339, request.NewStartTime)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid new start time"})
			return
		}
		newEndTime, err := time.Parse(time.RFC3339, request.NewEndTime)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid new end time"})
			return
		}
		appointment, err := appointmentService.RescheduleAppointment(uint(id), newStartTime, newEndTime)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.JSON(http.StatusOK, appointment)
	})

	router.DELETE("/appointments/:id", func(context *gin.Context) {
		idStr := context.Param("id")
		id, err := strconv.ParseUint(idStr, 10, 32)
		if err != nil {
			context.JSON(http.StatusBadRequest, gin.H{"error": "invalid appointment id"})
			return
		}
		if err := appointmentService.CancelAppointment(uint(id)); err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		context.Status(http.StatusNoContent)
	})

	if err := router.Run(":" + configuration.Port); err != nil {
		log.Fatalf("Failed to run the server: %v", err)
	}
}

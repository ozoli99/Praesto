package appointment

import (
	"errors"
	"fmt"
	"time"

	"github.com/ozoli99/Praesto/calendars"
	"github.com/ozoli99/Praesto/notifications"
)

type Service interface {
	BookAppointment(providerID, customerID uint, startTime, endTime time.Time) (*Appointment, error)
	RescheduleAppointment(appointmentID uint, newStartTime, newEndTime time.Time) (*Appointment, error)
	CancelAppointment(appointmentID uint) error
}

type AppointmentService struct {
	Repository                Repository
	NotificationService       notifications.NotificationService
	CalendarAdapter           calendars.CalendarAdapter
	NotificationConfiguration notifications.NotificationConfig
}

func NewService(repository Repository, notificationService notifications.NotificationService, calendarAdapter calendars.CalendarAdapter, notificationConfiguration notifications.NotificationConfig) Service {
	return &AppointmentService{
		Repository:                repository,
		NotificationService:       notificationService,
		CalendarAdapter:           calendarAdapter,
		NotificationConfiguration: notificationConfiguration,
	}
}

func (service *AppointmentService) BookAppointment(providerID, customerID uint, startTime, endTime time.Time) (*Appointment, error) {
	overlapping, err := service.Repository.FindOverlapping(providerID, startTime, endTime)
	if err != nil {
		return nil, err
	}
	if len(overlapping) > 0 {
		return nil, errors.New("the time slot conflicts with an existing appointment")
	}

	appointment := &Appointment{
		ProviderID: providerID,
		CustomerID: customerID,
		StartTime:  startTime,
		EndTime:    endTime,
		Status:     StatusBooked,
	}
	if err := service.Repository.Create(appointment); err != nil {
		return nil, err
	}

	if err := service.CalendarAdapter.SyncAppointment(appointment); err != nil {
		fmt.Printf("calendar sync error: %v\n", err)
	}
	service.NotificationService.ScheduleReminder(appointment, service.NotificationConfiguration)

	return appointment, nil
}

func (service *AppointmentService) RescheduleAppointment(appointmentID uint, newStartTime, newEndTime time.Time) (*Appointment, error) {
	appointment, err := service.Repository.GetByID(appointmentID)
	if err != nil {
		return nil, err
	}

	overlapping, err := service.Repository.FindOverlapping(appointment.ProviderID, newStartTime, newEndTime)
	if err != nil {
		return nil, err
	}
	for _, appt := range overlapping {
		if appt.ID != appointment.ID {
			return nil, fmt.Errorf("the new time slot conflicts with an existing appointment")
		}
	}

	appointment.StartTime = newStartTime
	appointment.EndTime = newEndTime
	appointment.Status = StatusRescheduled
	if err := service.Repository.Update(appointment); err != nil {
		return nil, err
	}

	if err := service.CalendarAdapter.SyncAppointment(appointment); err != nil {
		fmt.Printf("calendar sync error: %v\n", err)
	}
	service.NotificationService.ScheduleReminder(appointment, service.NotificationConfiguration)

	return appointment, nil
}

func (service *AppointmentService) CancelAppointment(appointmentID uint) error {
	appointment, err := service.Repository.GetByID(appointmentID)
	if err != nil {
		return err
	}
	appointment.Status = StatusCanceled
	if err := service.Repository.Update(appointment); err != nil {
		return err
	}

	if err := service.CalendarAdapter.RemoveAppointment(appointment); err != nil {
		fmt.Printf("calendar removal error: %v\n", err)
	}
	service.NotificationService.CancelReminder(appointment)

	return nil
}

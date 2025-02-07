package appointment

import (
	"errors"
	"fmt"
	"time"

	"github.com/ozoli99/Praesto/pkg/integration"
	"github.com/ozoli99/Praesto/pkg/notifications"
)

type Service interface {
	BookAppointment(providerID, customerID uint, startTime, endTime time.Time) (*Appointment, error)
	RescheduleAppointment(appointmentID uint, newStartTime, newEndTime time.Time) (*Appointment, error)
	CancelAppointment(appointmentID uint) error
}

type AppointmentService struct {
	Repository Repository
}

func NewService(repository Repository) Service {
	return &AppointmentService{Repository: repository}
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
	integration.SyncAppointmentToCalendar(appointment)
	notifications.ScheduleReminder(appointment)
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
	integration.SyncAppointmentToCalendar(appointment)
	notifications.ScheduleReminder(appointment)
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
	integration.RemoveAppointmentFromCalendar(appointment)
	notifications.CancelReminder(appointment)
	return nil
}
package appointment

import (
	"time"

	"github.com/ozoli99/Praesto/pkg/integration"
	"github.com/ozoli99/Praesto/pkg/notifications"
)

type Service interface {
	BookAppointment(providerID, customerID uint, startTime, endTime time.Time) (*Appointment, error)
	RescheduleAppointment(appointmentID uint, newStart, newEnd time.Time) (*Appointment, error)
	CancelAppointment(appointmentID uint) error
}

type AppointmentService struct {
	Repository Repository
}

func NewService(repository Repository) Service {
	return &AppointmentService{Repository: repository}
}

func (service *AppointmentService) BookAppointment(providerID, customerID uint, startTime, endTime time.Time) (*Appointment, error) {
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

func (service *AppointmentService) RescheduleAppointment(appointmentID uint, newStart, newEnd time.Time) (*Appointment, error) {
	appointment, err := service.Repository.GetByID(appointmentID)
	if err != nil {
		return nil, err
	}
	appointment.StartTime = newStart
	appointment.EndTime = newEnd
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
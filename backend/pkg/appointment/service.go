package appointment

type Service struct {
	Repository Repository
}

func NewService(repository Repository) *Service {
	return &Service{Repository: repository}
}

func (service *Service) BookAppointment(appointment *Appointment) error {
	appointment.Status = "booked"
	return service.Repository.Create(appointment)
}
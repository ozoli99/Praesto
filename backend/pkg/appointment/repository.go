package appointment

import "gorm.io/gorm"

type Repository interface {
	Create(appointment *Appointment) error
	GetByID(id uint) (*Appointment, error)
}

type GormRepository struct {
	Database *gorm.DB
}

func NewGormRepository(database *gorm.DB) *GormRepository {
	return &GormRepository{Database: database}
}

func (repository *GormRepository) Create(appointment *Appointment) error {
	return repository.Database.Create(appointment).Error
}

func (repository *GormRepository) GetByID(id uint) (*Appointment, error) {
	var appointment Appointment
	err := repository.Database.First(&appointment, id).Error
	return &appointment, err
}
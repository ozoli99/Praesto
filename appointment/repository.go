package appointment

import (
	"time"

	"gorm.io/gorm"
)

type Repository interface {
	Create(appointment *Appointment) error
	Update(appointment *Appointment) error
	Delete(id uint) error
	GetByID(id uint) (*Appointment, error)
	FindOverlapping(providerID uint, startTime, endTime time.Time) ([]Appointment, error)
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

func (repository *GormRepository) Update(appointment *Appointment) error {
	return repository.Database.Save(appointment).Error
}

func (repository *GormRepository) Delete(id uint) error {
	return repository.Database.Delete(&Appointment{}, id).Error
}

func (repository *GormRepository) GetByID(id uint) (*Appointment, error) {
	var appointment Appointment
	if err := repository.Database.First(&appointment, id).Error; err != nil {
		return nil, err
	}
	return &appointment, nil
}

func (repository *GormRepository) FindOverlapping(providerID uint, startTime, endTime time.Time) ([]Appointment, error) {
	var appointments []Appointment
	err := repository.Database.Where("provider_id = ? AND start_time < ? AND end_time > ?", providerID, endTime, startTime).Find(&appointments).Error
	return appointments, err
}
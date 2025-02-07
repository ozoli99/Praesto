package session

import "gorm.io/gorm"

type Repository interface {
	CreateSession(session *Session) error
	GetSessionByID(id uint) (*Session, error)
	UpdateSession(session *Session) error

	CreateTreatmentPlan(plan *TreatmentPlan) error
	GetTreatmentPlanByID(id uint) (*TreatmentPlan, error)
	UpdateTreatmentPlan(plan *TreatmentPlan) error
}

type GormRepository struct {
	Database *gorm.DB
}

func NewGormRepository(database *gorm.DB) Repository {
	return &GormRepository{Database: database}
}

func (repository *GormRepository) CreateSession(session *Session) error {
	return repository.Database.Create(session).Error
}

func (repository *GormRepository) GetSessionByID(id uint) (*Session, error) {
	var session Session
	if err := repository.Database.First(&session, id).Error; err != nil {
		return nil, err
	}
	return &session, nil
}

func (repository *GormRepository) UpdateSession(session *Session) error {
	return repository.Database.Save(session).Error
}

func (repository *GormRepository) CreateTreatmentPlan(plan *TreatmentPlan) error {
	return repository.Database.Create(plan).Error
}

func (repository *GormRepository) GetTreatmentPlanByID(id uint) (*TreatmentPlan, error) {
	var plan TreatmentPlan
	if err := repository.Database.First(&plan, id).Error; err != nil {
		return nil, err
	}
	return &plan, nil
}

func (repository *GormRepository) UpdateTreatmentPlan(plan *TreatmentPlan) error {
	return repository.Database.Save(plan).Error
}

package repository

import (
	"github.com/ozoli99/Praesto/pkg/models"
	"gorm.io/gorm"
)

type UserRepository interface {
	Create(user *models.User) error
	GetByID(id uint) (*models.User, error)
	Update(user *models.User) error
	Delete(id uint) error
}

type GormUserRepository struct {
	Database *gorm.DB
}

func NewGormUserRepository(database *gorm.DB) *GormUserRepository {
	return &GormUserRepository{Database: database}
}

func (repository *GormUserRepository) Create(user *models.User) error {
	return repository.Database.Create(user).Error
}

func (repository *GormUserRepository) GetByID(id uint) (*models.User, error) {
	var user models.User
	err := repository.Database.First(&user, id).Error
	return &user, err
}

func (repository *GormUserRepository) Update(user *models.User) error {
	return repository.Database.Save(user).Error
}

func (repository *GormUserRepository) Delete(id uint) error {
	return repository.Database.Delete(&models.User{}, id).Error
}
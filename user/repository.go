package user

import (
	"gorm.io/gorm"
)

type Repository interface {
	Create(user *User) error
	GetByAuth0ID(auth0ID string) (*User, error)
	GetByID(id uint) (*User, error)
	Update(user *User) error
	Delete(id uint) error
}

type GormRepository struct {
	Database *gorm.DB
}

func NewGormRepository(database *gorm.DB) *GormRepository {
	return &GormRepository{Database: database}
}

func (repository *GormRepository) Create(user *User) error {
	return repository.Database.Create(user).Error
}

func (repository *GormRepository) GetByAuth0ID(auth0ID string) (*User, error) {
	var user User
	if err := repository.Database.Where("auth0_id = ?", auth0ID).First(&user).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *GormRepository) GetByID(id uint) (*User, error) {
	var user User
	if err := repository.Database.First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}

func (repository *GormRepository) Update(user *User) error {
	return repository.Database.Save(user).Error
}

func (repository *GormRepository) Delete(id uint) error {
	return repository.Database.Delete(&User{}, id).Error
}
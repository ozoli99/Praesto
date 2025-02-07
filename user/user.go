package user

import (
	"time"

	"github.com/ozoli99/Praesto/models"
)

const (
	RoleClient  = "client"
	RoleMasseur = "masseur"
	RoleAdmin   = "admin"
)

type User struct {
	models.Base
	Auth0ID        string `gorm:"size:255;uniqueIndex"`
	Name           string `gorm:"size:255"`
	Email          string `gorm:"size:255;uniqueIndex"`
	ProfilePicture string `gorm:"size:512"`
	Role           string `gorm:"size:50"`
	Verified       bool
	LastLogin      time.Time
}

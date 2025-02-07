package user

import (
	"time"
)

type Service interface {
	SyncUserFromClaims(claims map[string]interface{}) (*User, error)
	UpdateProfile(user *User) error
}

type UserService struct {
	repository Repository
}

func NewService(repository Repository) Service {
	return &UserService{repository: repository}
}

func (service *UserService) SyncUserFromClaims(claims map[string]interface{}) (*User, error) {
	sub, _ := claims["sub"].(string)
	email, _ := claims["email"].(string)
	name, _ := claims["name"].(string)
	picture, _ := claims["picture"].(string)

	var role string
	namespace := "https://yourdomain.com/"
	if roles, ok := claims[namespace+"roles"].([]interface{}); ok && len(roles) > 0 {
		role, _ = roles[0].(string)
	} else {
		role = RoleClient
	}

	user, err := service.repository.GetByAuth0ID(sub)
	if err != nil {
		newUser := &User{
			Auth0ID:        sub,
			Email:          email,
			Name:           name,
			ProfilePicture: picture,
			Role:           role,
			Verified:       true,
			LastLogin:      time.Now(),
		}
		if err := service.repository.Create(newUser); err != nil {
			return nil, err
		}
		return newUser, nil
	}

	user.Email = email
	user.Name = name
	user.ProfilePicture = picture
	user.LastLogin = time.Now()
	if err := service.repository.Update(user); err != nil {
		return nil, err
	}
	return user, nil
}

func (service *UserService) UpdateProfile(user *User) error {
	return service.repository.Update(user)
}

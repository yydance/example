package service

import (
	"demo-base/internal/models"
	"errors"
)

// Service is the interface for the service layer
// It should be implemented by any service that needs to interact with the database

type UserInput struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Role  string `json:"role"`
}

type UserService interface {
	Create() error
	Get() error
	Update() error
	Delete() error
	List(pageNum, pageSize int) error
}

func (u *UserInput) Create() error {
	if u.IsExist() {
		return errors.New("user already exists")
	}
	user := models.User{
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role,
	}
	return user.Create()
}

func (u *UserInput) Get() (models.User, error) {
	user := models.User{
		Name: u.Name,
	}
	return user, user.Find()
}

func (u *UserInput) Update() error {
	if u.IsExist() {
		return errors.New("user already exists")
	}
	user := models.User{
		Name:  u.Name,
		Email: u.Email,
		Role:  u.Role,
	}
	return user.Update()

}

func (u *UserInput) Delete() error {
	user := models.User{
		Name: u.Name,
	}
	return user.Delete()
}

func (u *UserInput) List(pageNum, pageSize int) ([]models.User, error) {
	var user models.User
	return user.List(pageNum, pageSize)

}

func (u *UserInput) IsExist() bool {
	var user models.User
	return user.IsExist(u.Name)
}

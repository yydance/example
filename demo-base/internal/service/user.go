package service

import (
	"demo-base/internal/models"
	"errors"
)

// Service is the interface for the service layer
// It should be implemented by any service that needs to interact with the database

type UserInput struct {
	Name         string   `json:"name" validate:"required,max=64"`
	Email        string   `json:"email" validate:"omitempty,email"`
	RolePlatform []string `json:"role_platform" validate:"omitempty"`
	Password     string   `json:"password" validate:"omitempty,min=6,max=64,alphanum"`
}

type UserPassword struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=64,alphanum"`
}

type UsersOutput struct {
	Name         string   `json:"name"`
	Email        string   `json:"email"`
	RolePlatform []string `json:"role_platform"`
	CreatedAt    string   `json:"created_at"`
	UpdatedAt    string   `json:"updated_at"`
}

/*
type UserService interface {
	Create() error
	Get() (models.User, error)
	Update() error
	Delete() error
	List(pageNum, pageSize int) ([]models.User, error)
}
*/

func (u *UserInput) Create() error {
	err := validate.Struct(u)
	if err != nil || u.Password == "" {
		return errors.New("invalid input")
	}
	if u.IsExist() {
		return errors.New("user already exists")
	}
	user := models.User{
		Name:         u.Name,
		Email:        u.Email,
		Password:     u.Password,
		RolePlatform: u.RolePlatform,
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
	err := validate.Struct(u)
	if err != nil {
		return errors.New("invalid input")
	}
	if u.IsExist() {
		return errors.New("user already exists")
	}
	user := models.User{
		Name:         u.Name,
		Email:        u.Email,
		RolePlatform: u.RolePlatform,
	}
	return user.Update()

}

func (u *UserPassword) UpdatePassword() error {
	user := models.User{
		Name:     u.Name,
		Password: u.Password,
	}
	return user.UpdatePassword()
}

func (u *UserInput) Delete() error {
	user := models.User{
		Name: u.Name,
	}
	return user.Delete()
}

func (u *UserInput) List(pageNum, pageSize int) ([]UsersOutput, error) {
	var user models.User
	users := make([]UsersOutput, 0)
	userList, err := user.List(pageNum, pageSize)
	if err != nil {
		return users, err
	}
	for _, user := range userList {
		users = append(users, UsersOutput{
			Name:         user.Name,
			Email:        user.Email,
			RolePlatform: user.RolePlatform,
			CreatedAt:    user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:    user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return users, nil
}

func (u *UserInput) IsExist() bool {
	var user models.User
	return user.IsExist(u.Name)
}

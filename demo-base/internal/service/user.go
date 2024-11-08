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
	if _, ok := u.IsExist(); ok {
		return errors.New("user already exists")
	}
	user := models.User{
		Name:         u.Name,
		Email:        u.Email,
		Password:     u.Password,
		RolePlatform: u.RolePlatform,
	}

	if err := user.Create(); err == nil {
		var rbac *RBACPlatform
		if err := rbac.AddGroupPolicies(u.Name, u.RolePlatform); err != nil {
			return errors.New("Failed to add group policies")
		}
	} else {
		return err
	}
	return nil
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
	if _, ok := u.IsExist(); ok {
		return errors.New("user already exists")
	}
	user := models.User{
		Name:         u.Name,
		Email:        u.Email,
		RolePlatform: u.RolePlatform,
	}
	if err := user.Update(); err == nil {
		var rbac *RBACPlatform
		if err := rbac.UpdateGroupPolicies(u.Name, u.RolePlatform); err != nil {
			return errors.New("Failed to update group policies")
		}
	} else {
		return err
	}
	return nil
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
	if err := user.Delete(); err == nil {
		var rbac *RBACPlatform
		if err := rbac.RemoveGroupPolicies(u.Name); err != nil {
			return errors.New("Failed to delete group policies")
		}
	} else {
		return err
	}
	return nil
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

func (u *UserInput) IsExist() (models.User, bool) {
	var user models.User
	return user.IsExist(u.Name)
}

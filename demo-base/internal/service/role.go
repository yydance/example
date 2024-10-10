package service

import (
	"demo-base/internal/models"
	"errors"
)

type RoleInput struct {
	Name        string   `json:"name"`
	Description string   `json:"description"`
	Permissions []string `json:"permissions"`
}

func (r *RoleInput) Create() error {
	role := models.Role{
		Name:        r.Name,
		Description: r.Description,
		Permissions: r.Permissions,
	}
	if r.IsExist() {
		return errors.New("role already exists")
	}

	return role.Create()
}

func (r *RoleInput) Update() error {
	role := models.Role{
		Name:        r.Name,
		Description: r.Description,
		Permissions: r.Permissions,
	}
	if r.IsExist() {
		return errors.New("role already exists")
	}

	return role.Update()
}

func (r *RoleInput) IsExist() bool {
	var role models.Role
	return role.IsExist(r.Name)
}

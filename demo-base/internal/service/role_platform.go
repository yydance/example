package service

import (
	"demo-base/internal/models"
	"errors"
)

type RolePlatform struct {
	Name        string   `json:"name" validate:"required,max=64"`
	Description string   `json:"description" validate:"max=256"`
	Permissions []string `json:"permissions"`
}

func (r *RolePlatform) Create() error {
	err := validate.Struct(r)
	if err != nil {
		return errors.New("invalid input")
	}

	role := models.RolePlatform{
		Name:        r.Name,
		Description: r.Description,
		Permissions: r.Permissions,
	}

	return role.Create()
}

func (r *RolePlatform) Update() error {
	err := validate.Struct(r)
	if err != nil {
		return errors.New("invalid input")
	}

	role := models.RolePlatform{
		Name:        r.Name,
		Description: r.Description,
		Permissions: r.Permissions,
	}
	return role.Update()
}

func (r *RolePlatform) Delete() error {

	role := models.RolePlatform{
		Name: r.Name,
	}
	// rollback，删除失败，需要恢复rbac中的数据，先查询出原来的数据，再进行恢复

	return role.Delete()
}

func (r *RolePlatform) GetRoles() ([]string, error) {
	role := models.RolePlatform{
		Name: r.Name,
	}
	return role.Permissions, role.Find()
}

/*
func (r *RolePlatform) IsExists() bool {
	role := models.RolePlatform{
		Name: r.Name,
	}
	return role.IsExist()
}
*/

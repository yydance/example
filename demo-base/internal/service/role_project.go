package service

import (
	"demo-base/internal/models"
	"errors"
)

type RoleProject struct {
	Name        string   `json:"name" validate:"required,max=64,alphanum"`
	Permissions []string `json:"permissions" validate:"required"`
}

func (r *RoleProject) Create(project_name string) error {
	err := validate.Struct(r)
	if err != nil || project_name == "" {
		return errors.New("参数错误")
	}
	role := models.RoleProject{
		RoleName:    r.Name,
		Permissions: r.Permissions,
		ProjectName: project_name,
	}
	return role.Create()
}

func (r *RoleProject) Update(project_name string) error {
	err := validate.Struct(r)
	if err != nil || project_name == "" {
		return errors.New("参数错误")
	}
	role := models.RoleProject{
		RoleName:    r.Name,
		Permissions: r.Permissions,
		ProjectName: project_name,
	}
	return role.Update()
}

func (r *RoleProject) Delete(project_name string) error {
	if project_name == "" {
		return errors.New("缺少项目名称")
	}
	role := models.RoleProject{
		RoleName:    r.Name,
		ProjectName: project_name,
	}
	return role.Delete()
}

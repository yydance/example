package models

import "errors"

type RoleProject struct {
	ProjectName string   `json:"project_name" gorm:"type:varchar(64);not null"`
	RoleName    string   `json:"name" gorm:"type:varchar(64);not null"`
	Permissions []string `json:"permissions" gorm:"type:json;not null"`
}

func (RoleProject) TableName() string {
	return "roles_project"
}

func (r *RoleProject) Create() error {
	if r.IsExist() {
		return errors.New("项目角色已存在")
	}
	return DB.Create(r).Error
}

func (r *RoleProject) Update() error {
	return DB.Where("role_name = ? and project_name = ?", r.RoleName, r.ProjectName).Updates(r).Error
}

func (r *RoleProject) Delete() error {
	return DB.Where("role_name = ? and project_name = ?", r.RoleName, r.ProjectName).Delete(r).Error
}

func (r *RoleProject) List() ([]RoleProject, error) {
	var roles []RoleProject
	return roles, DB.Find(&roles).Error
}

func (r *RoleProject) IsExist() bool {
	var role RoleProject
	return DB.Where("role_name = ? and project_name = ?", r.RoleName, r.ProjectName).First(&role).RowsAffected > 0
}

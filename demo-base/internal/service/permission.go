package service

import "demo-base/internal/models"

func ListRolePermission() models.RolePermission {
	var RolePermission *models.RolePermission
	return RolePermission.NewRolePermissions()
}

func ListProjectPermission() models.ProjectPermission {
	var ProjectPermission *models.ProjectPermission
	return ProjectPermission.NewProjectPermissions()
}

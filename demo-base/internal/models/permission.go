package models

/*
type Permission struct {
	Id   int64  `json:"id" gorm:"primarykey"`
	Name string `json:"name"`
}

func (Permission) TableName() string {
	return "permissions"
}

func (p *Permission) List() ([]Permission, error) {
	var ps []Permission
	return ps, DB.Find(&ps).Error
}

func (p *Permission) Create() error {
	return DB.Create(p).Error
}
*/

type RolePermission []string

func (r RolePermission) NewRolePermissions() RolePermission {
	r = []string{"User Management", "Role Management", "Permission Management"}
	return r
}

type ProjectPermission []string

func (p ProjectPermission) NewProjectPermissions() ProjectPermission {
	p = []string{"Project View", "Role Management", "Project Admin"}
	return p
}

package models

import "errors"

type RouteUri string

type RolePlatform struct {
	Name        string   `json:"name" gorm:"type:varchar(64);not null"`
	Description string   `json:"description" gorm:"type:varchar(255)"`
	Permissions []string `json:"permissions" gorm:"type:json"`
}

func (RolePlatform) TableName() string {
	return "roles_platform"
}

func (r *RolePlatform) Create() error {
	if r.IsExist() {
		return errors.New("平台角色已存在")
	}
	return DB.Create(r).Error
}
func (r *RolePlatform) Update() error {
	return DB.Where("name = ?", r.Name).Updates(r).Error
}
func (r *RolePlatform) Delete() error {
	return DB.Where("name = ?", r.Name).Delete(r).Error
}

func (r *RolePlatform) List(pageNum, pageLimit int) ([]RolePlatform, error) {
	var roles []RolePlatform
	return roles, DB.Offset((pageNum - 1) * pageLimit).Limit(pageLimit).Find(&roles).Error
}

func (r *RolePlatform) Find() error {
	return DB.First(r, r.Name).Error
}

var permissions map[string]RouteUri = map[string]RouteUri{
	"User Management": "/api/v1/users",
	"Role Management": "/api/v1/roles",
}

func ListPermissions() []string {

	keys := make([]string, 0, len(permissions))
	for k := range permissions {
		keys = append(keys, k)
	}
	return keys
}

func (r *RolePlatform) IsExist() bool {
	var role RolePlatform
	return DB.Where("name = ?", r.Name).First(&role).RowsAffected > 0
}

package models

import "gorm.io/gorm"

type RouteUri string

type Role struct {
	gorm.Model
	Name        string   `json:"name" gorm:"type:varchar(64);not null"`
	Description string   `json:"description" gorm:"type:varchar(255)"`
	Permissions []string `json:"permissions" gorm:"type:json"`
}

func (Role) TableName() string {
	return "roles"
}

func (r *Role) Create() error {
	return DB.Create(r).Error
}
func (r *Role) Update() error {
	return DB.Where("name = ?", r.Name).Save(r).Error
}
func (r *Role) Delete() error {
	return DB.Where("name = ?", r.Name).Delete(r).Error
}

func (r *Role) List(pageNum, pageLimit int) ([]Role, error) {
	var roles []Role
	return roles, DB.Offset((pageNum - 1) * pageLimit).Limit(pageLimit).Find(&roles).Error
}

func (r *Role) Find() error {
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

func (r *Role) IsExist(name string) bool {
	var role Role
	return DB.Where("name = ?", name).First(&role).Error == nil
}

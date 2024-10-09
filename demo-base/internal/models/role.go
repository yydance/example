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
	// TODO: Implement role creation logic
	return DB.Create(r).Error
}
func (r *Role) Update() error {
	// TODO: Implement role update logic
	return DB.Save(r).Error
}
func (r *Role) Delete() error {
	// TODO: Implement role deletion logic
	return DB.Delete(r).Error
}

func (r *Role) List() ([]Role, error) {
	// TODO: Implement role listing logic
	var roles []Role
	return roles, DB.Find(&roles).Error
}

func (r *Role) Get(id uint) error {
	// TODO: Implement role retrieval logic
	return DB.First(r, id).Error
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

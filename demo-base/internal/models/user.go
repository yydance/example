package models

import (
	"errors"

	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string   `json:"name" gorm:"type:varchar(50);not null"`
	Email    string   `json:"email" gorm:"type:varchar(50)"`
	Password string   `json:"password" gorm:"type:varchar(50)"`
	Roles    []string `json:"roles" gorm:"column:roles;serializer:json"`
	//RoleProject  string   `json:"role_project_id" gorm:"column:role_project_id;type:varchar(50)"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Create() error {
	if _, ok := u.IsExist(); ok {
		return errors.New("user already exists")
	}
	return DB.Create(u).Error
}

func (u *User) Update() error {
	if _, ok := u.IsExist(); !ok {
		return errors.New("user does not exist")
	}
	return DB.Where("name = ?", u.Name).Updates(u).Error
}

func (u *User) UpdatePassword() error {
	return DB.Where("name = ?", u.Name).Update("password", u.Password).Error
}

func (u *User) Delete() error {
	return DB.Where("name = ?", u.Name).Delete(u).Error
}

func (u *User) Find() error {
	//return DB.First(u, u.Name).Error
	return DB.Where("name = ?", u.Name).First(u).Error
}
func (u *User) FindAllByName() ([]string, error) {
	var names []string
	return names, DB.Select([]string{"name"}).Scan(&names).Error
}

func (u *User) FindByName(name string) error {
	return DB.Where("name = ?", name).First(u).Error
}

func (u *User) FindByEmail(email string) error {
	return DB.Where("email = ?", email).First(u).Error
}

func (u *User) List(pageNum, pageLimit int) ([]User, error) {
	var users []User
	return users, DB.Offset((pageNum - 1) * pageLimit).Limit(pageLimit).Find(&users).Error
}

func (u *User) Count() (int64, error) {
	var count int64
	return count, DB.Model(&User{}).Count(&count).Error
}

func (u *User) IsExist() (User, bool) {
	var user User
	return user, DB.Where("name = ?", u.Name).First(&user).RowsAffected > 0
}

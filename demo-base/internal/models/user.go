package models

import (
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name  string `json:"name" gorm:"type:varchar(50);not null"`
	Email string `json:"email" gorm:"type:varchar(50)"`
	Role  string `json:"role_name" gorm:"column:role_name;type:varchar(50)"`
}

func (User) TableName() string {
	return "users"
}

func (u *User) Create() error {
	return DB.Create(u).Error
}

func (u *User) Update() error {
	return DB.Save(u).Error
}

func (u *User) Delete() error {
	return DB.Where("id = ?", u.ID).Delete(u).Error
}

func (u *User) Find() error {
	return DB.First(u, u.ID).Error
}

func (u *User) FindByName(name string) error {
	return DB.Where("name = ?", name).First(u).Error
}

func (u *User) FindByEmail(email string) error {
	return DB.Where("email = ?", email).First(u).Error
}

func (u *User) FindAll() ([]User, error) {
	var users []User
	return users, DB.Find(&users).Error
}

func (u *User) FindByPage(pageNum, pageLimit int) ([]User, error) {
	var users []User
	return users, DB.Offset((pageNum - 1) * pageLimit).Limit(pageLimit).Find(&users).Error
}

func (u *User) Count() (int64, error) {
	var count int64
	return count, DB.Model(&User{}).Count(&count).Error
}

func (u *User) IsExist(name string) (bool, error) {
	var user User
	return user.ID != 0, DB.Where("name = ?", name).First(&user).Error
}

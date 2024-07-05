package models

import (
	"time"

	"gorm.io/gorm"
)

func UserList(pageSize, pageNum int) (any, error) {
	tx := initCtx(5 * time.Second)

	var data []*User
	err := tx.Where(&User{}).Offset(pageNum).Limit(pageSize).Find(data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return data, nil
}

func (u User) Add() error {

	return nil
}

func (u User) UserDelete() error {

	return nil
}

func UserBatchDelete() error {

	return nil
}

func UserEdit(id int) error {
	return nil
}

/*
func (u User) Info() (any, error) {
	CloseDB()
	tx := initCtx(5 * time.Second)
	var data User
	if strconv.Itoa(int(u.ID)) != "" {
		err := tx.Where("id = ?", u.ID).First(&data).Error
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	if u.Username != "" {
		err := tx.Where("username = ?", u.Username).First(&data).Error
		if err != nil {
			return nil, err
		}
		return data, nil
	}
	return data, nil
}
*/

func UserToken() (any, error) {

	return nil, nil
}

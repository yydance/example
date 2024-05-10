package models

import (
	"demo-dashboard/internal/log"
	"fmt"
	"time"

	"gorm.io/gorm"
)

func CreateUpstream(id, name, upstream_type, desc string, create_time, update_time int64) error {
	CloseDB()
	tx := initCtx(10 * time.Second)

	upstream := Upstream{
		Name: name,
		Type: upstream_type,
		Desc: desc,
		BaseInfo: BaseInfo{
			ID:         id,
			CreateTime: create_time,
			UpdateTime: update_time,
		},
	}
	res := tx.Model(&Upstream{}).Create(&upstream)
	if res.Error != nil {
		log.Logger.Error(res.Error)
		return res.Error
	}
	return nil

}

func GetUpstreamList(pageNum int, pageSize int) (any, error) {
	CloseDB()
	tx := initCtx(10 * time.Second)

	var data []*Upstream
	var maps Upstream
	err := tx.Where(&maps).Offset(pageNum).Limit(pageSize).Find(data).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		return nil, err
	}
	return data, nil
}

func UpdateUpstream(id int) error {
	CloseDB()

	tx := initCtx(5 * time.Second)
	var data []*Upstream
	err := tx.Model(&Upstream{}).Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUpstreamByID(id string) (any, error) {
	CloseDB()
	tx := initCtx(5 * time.Second)
	var data Upstream
	err := tx.Where("id = ?", id).First(&data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

func GetUpstreamByName(name string) (any, error) {
	CloseDB()
	tx := initCtx(5 * time.Second)
	var data []*Upstream
	err := tx.Where("name LIKE ?", fmt.Sprintf("%%%s%%", name)).Find(data).Error
	if err != nil {
		return nil, err
	}
	return data, nil
}

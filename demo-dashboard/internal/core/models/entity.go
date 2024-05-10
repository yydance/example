package models

import "gorm.io/gorm"

type BaseInfo struct {
	ID         any   `json:"id" gorm:"primaryKey,uniqueIndex:idx_id"`
	CreateTime int64 `json:"create_time,omitempty"`
	UpdateTime int64 `json:"update_time,omitempty"`
}

type Upstream struct {
	BaseInfo
	Name string `json:"name,omitempty"`
	Type string `json:"type,omitempty"`
	Desc string `json:"desc,omitempty"`
}

type Service struct {
	BaseInfo
	Name string `json:"type,omitempty"`
	Desc string `json:"desc,omitempty"`
}

type Route struct {
	BaseInfo
	Name   string   `json:"type,omitempty"`
	Hosts  []string `json:"hosts,omitempty"`
	Uris   []string `json:"uris,omitempty"`
	Status bool     `json:"status"`
}

type User struct {
	gorm.Model
	Username string `gorm:"username"`
	Password string `gorm:"password"`
	Email    string `gorm:"email"`
	Project  string `gorm:"project"`
}

type Project struct {
	gorm.Model
	Name string `gorm:"name"`
}

package models

type BaseInfo struct {
	ID         any   `json:"id" gorm:"primaryKey"`
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

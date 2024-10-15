package models

import (
	"errors"

	"gorm.io/gorm"
)

type ProjectBase struct {
	gorm.Model
	Name        string `json:"name" gorm:"max=128;not null;comment:项目名称"`
	Admin       string `json:"admin" gorm:"max=64;not null;comment:项目管理员"`
	Description string `json:"description" gorm:"type:text;comment:项目描述"`
}

func (ProjectBase) TableName() string {
	return "projects_base"
}

func (p *ProjectBase) Create() error {
	if p.IsExist() {
		return errors.New("项目已存在")
	}
	return DB.Create(p).Error
}

func (p *ProjectBase) Update() error {
	return DB.Where("name = ?", p.Name).Updates(p).Error
}

func (p *ProjectBase) Delete() error {
	return DB.Where("name = ?", p.Name).Delete(p).Error
}

func (p *ProjectBase) List(pageNum, pageSize int) ([]ProjectBase, error) {
	var projects []ProjectBase
	return projects, DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&projects).Error
}

func (p *ProjectBase) IsExist() bool {
	var project ProjectBase
	return DB.Where("name = ?", p.Name).First(&project).RowsAffected > 0
}

type ProjectMember struct {
	gorm.Model
	ProjectName string   `json:"project_name" gorm:"max=64;not null;comment:项目名称"`
	UserName    string   `json:"user_name" gorm:"max=64;not null;comment:用户名称"`
	Role        []string `json:"role" gorm:"type:json;not null;comment:项目角色"`
}

func (ProjectMember) TableName() string {
	return "projects_member"
}

func (p *ProjectMember) Add() error {
	if p.IsExist() {
		return errors.New("项目成员已存在")
	}
	return DB.Create(p).Error
}

func (p *ProjectMember) Update() error {
	return DB.Where("project_name = ? and user_name = ?", p.ProjectName, p.UserName).Updates(p).Error
}

func (p *ProjectMember) Delete() error {
	return DB.Where("project_name = ? and user_name = ?", p.ProjectName, p.UserName).Delete(p).Error
}
func (p *ProjectMember) List(pageNum, pageSize int) ([]ProjectMember, error) {
	var projectMembers []ProjectMember
	return projectMembers, DB.Offset((pageNum - 1) * pageSize).Limit(pageSize).Find(&projectMembers).Error
}

func (p *ProjectMember) IsExist() bool {
	var projectMember ProjectMember
	return DB.Where("project_name = ? and user_name = ?", p.ProjectName, p.UserName).First(&projectMember).RowsAffected > 0
}

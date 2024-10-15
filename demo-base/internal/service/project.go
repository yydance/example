package service

import (
	"demo-base/internal/models"
	"errors"
)

type ProjectsOutput struct {
	Name        string `json:"name"`
	Admin       string `json:"admin"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type ProjectInput struct {
	Name        string `json:"name" validate:"required,max=64,alphanum"`
	Admin       string `json:"admin" validate:"required"`
	Description string `json:"description" validate:"max=256"`
}

func (p *ProjectInput) Create() error {
	err := validate.Struct(p)
	if err != nil {
		return errors.New("参数错误")
	}
	project := models.ProjectBase{
		Name:        p.Name,
		Admin:       p.Admin,
		Description: p.Description,
	}
	return project.Create()
}

func (p *ProjectInput) Update() error {
	err := validate.Struct(p)
	if err != nil {
		return errors.New("参数错误")
	}
	project := models.ProjectBase{
		Name:        p.Name,
		Admin:       p.Admin,
		Description: p.Description,
	}
	return project.Update()
}

func (p *ProjectInput) Delete() error {
	project := models.ProjectBase{
		Name: p.Name,
	}
	return project.Delete()
}

func (p *ProjectInput) List(pageNum, pageSize int) ([]ProjectsOutput, error) {
	var project models.ProjectBase
	projects_output := make([]ProjectsOutput, 0)
	projects, err := project.List(pageNum, pageSize)
	if err != nil {
		return projects_output, err
	}
	for _, project := range projects {
		projects_output = append(projects_output, ProjectsOutput{
			Name:        project.Name,
			Admin:       project.Admin,
			Description: project.Description,
			CreatedAt:   project.Model.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt:   project.Model.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return projects_output, nil
}

type ProjectMemberInput struct {
	UserName string   `json:"user_name" validate:"required"`
	Roles    []string `json:"roles" validate:"required"`
}

type ProjectMemberOutput struct {
	UserName  string   `json:"user_name"`
	Roles     []string `json:"roles"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

func (p *ProjectMemberInput) AddMember(project_name string) error {
	err := validate.Struct(p)
	if err != nil || project_name == "" {
		return errors.New("参数错误")
	}
	project_member := models.ProjectMember{
		ProjectName: project_name,
		UserName:    p.UserName,
		Role:        p.Roles,
	}
	return project_member.Add()
}

func (p *ProjectMemberInput) UpdateMember(project_name string) error {
	if p.UserName == "" || project_name == "" {
		return errors.New("参数错误")
	}
	project_member := models.ProjectMember{
		ProjectName: project_name,
		UserName:    p.UserName,
		Role:        p.Roles,
	}
	return project_member.Update()
}

func (p *ProjectMemberInput) DeleteMember(project_name string) error {
	if p.UserName == "" || project_name == "" {
		return errors.New("参数错误")
	}
	project_member := models.ProjectMember{
		ProjectName: project_name,
		UserName:    p.UserName,
	}
	return project_member.Delete()
}

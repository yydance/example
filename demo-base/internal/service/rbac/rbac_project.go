package rbac

import (
	"demo-base/internal/conf"
	"demo-base/internal/models"
	"demo-base/internal/utils/logger"
	"errors"
	"fmt"
	"os"

	"github.com/bytedance/sonic"
)

type RBACProject struct {
	Roles map[string]map[string]string `json:"roles"`
}

func NewRBACProject() *RBACProject {
	input, err := os.ReadFile(conf.RBACProjectPolicy)
	if err != nil {
		logger.Panic(err.Error())
	}
	var rbac *RBACProject
	if err := sonic.Unmarshal(input, &rbac.Roles); err != nil {
		logger.Panic(err.Error())
	}
	return rbac
}

func (r *RBACProject) AddPolicies(role string, project string, permissions []string) error {
	// TODO: check role and project

	res, err := models.CasbinEnforcer.AddPoliciesEx(r.GetPoliciesForRole(role, project, permissions))
	if err != nil || !res {
		return errors.New(fmt.Sprintf("Add Policy for project(%s) role(%s) failed: %v", project, role, err))
	}
	return nil
}

func (r *RBACProject) RemovePolicies(role string, project string) error {
	users := models.CasbinEnforcer.GetUsersForRoleInDomain(role, project)
	if len(users) > 0 {
		return errors.New(fmt.Sprintf("Project(%s) role(%s) has users, please remove users first", project, role))
	}
	res, err := models.CasbinEnforcer.DeleteRolesForUserInDomain(role, project)
	if err != nil {
		return err
	}
	if !res {
		return errors.New(fmt.Sprintf("Policies not fount for project(%s) role: %s", project, role))
	}
	return nil
}

func (r *RBACProject) UpdatePolicies(role string, project string, permissions []string) error {

	return nil
}

func (r *RBACProject) GetPoliciesForRole(role string, project string, permissions []string) [][]string {
	res := make([][]string, 0)
	for _, permission := range permissions {
		for path, methods := range r.Roles[permission] {
			res = append(res, []string{role, project, path, methods})
		}
	}
	return res
}

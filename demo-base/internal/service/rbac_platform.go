package service

import (
	"demo-base/internal/conf"
	"demo-base/internal/models"
	"demo-base/internal/utils/logger"
	"demo-base/internal/utils/tools"
	"errors"
	"fmt"
	"os"

	"github.com/bytedance/sonic"
)

// TODO: rbac policy schema validation

type RBACPlatform struct {
	Roles map[string]map[string]string `json:"roles"`
}

func NewRBACPlatform() *RBACPlatform {
	input, err := os.ReadFile(conf.RBACPlatformPolicy)
	if err != nil {
		logger.Panic(err.Error())
	}
	var rbac *RBACPlatform

	if err = sonic.Unmarshal(input, &rbac.Roles); err != nil {
		logger.Panic(err.Error())
	}
	return rbac
}

// rbac 角色策略，p, sub, dom, obj, act
func (r *RBACPlatform) AddPolicies(role string, permissions []string) error {
	// TODO: check role

	res, err := models.CasbinEnforcer.AddPoliciesEx(r.GetPoliciesForRole(role, permissions))
	if err != nil || !res {
		return errors.New(fmt.Sprintf("Add Policy for role(%s) failed: %v", role, err))
	}
	return nil
}

func (r *RBACPlatform) UpdatePolicies(role string, permissions []string) error {
	oldDiffPolicies, newDiffPolicies, err := r.GetDiffPolicies(role, permissions)
	if err != nil {
		return err
	}
	if len(oldDiffPolicies) > 0 {
		res, err := models.CasbinEnforcer.RemovePolicies(oldDiffPolicies)
		if err != nil || !res {
			return errors.New(fmt.Sprintf("Remove Policy for role(%s) failed: %v", role, err))
		}
	}
	if len(newDiffPolicies) > 0 {
		res, err := models.CasbinEnforcer.AddPoliciesEx(newDiffPolicies)
		if err != nil || !res {
			return errors.New(fmt.Sprintf("Add Policy for role(%s) failed: %v", role, err))
		}
	}
	return nil
}

func (r *RBACPlatform) GetPolicies(role string) ([][]string, error) {
	return models.CasbinEnforcer.GetFilteredNamedPolicy("p", 0, role, "global")
}

func (r *RBACPlatform) RemovePolicies(role string) error {
	users := models.CasbinEnforcer.GetUsersForRoleInDomain(role, "global")
	if len(users) > 0 {
		return errors.New(fmt.Sprintf("Global role(%s) has users, please remove users first", role))
	}
	res, err := models.CasbinEnforcer.DeleteRolesForUserInDomain(role, "global")
	if err != nil {
		return err
	}
	if !res {
		return errors.New(fmt.Sprintf("Policies not fount for global role: %s", role))
	}
	return nil
}

func (r *RBACPlatform) AddPoliciesEx(rules [][]string) error {
	res, err := models.CasbinEnforcer.AddPoliciesEx(rules)
	if err != nil || !res {
		return errors.New("AddPolicies failed")
	}
	return nil
}

func (r *RBACPlatform) GetPoliciesForRole(role string, permissions []string) [][]string {
	res := make([][]string, 0)
	for _, permission := range permissions {
		for path, methods := range r.Roles[permission] {
			res = append(res, []string{role, "global", path, methods})
		}
	}
	return res
}

func (r *RBACPlatform) GetDiffPolicies(role string, permissions []string) ([][]string, [][]string, error) {
	oldDiffPolicies, newDiffPolicies := make([][]string, 0), make([][]string, 0)
	existPolicies := make([][]string, 0)
	oldPolicies, err := r.GetPolicies(role)
	//filteredPolicy := models.CasbinEnforcer.
	if err != nil {
		return nil, nil, err
	}
	for _, permission := range permissions {
		objects := r.Roles[permission]
		for path, methods := range objects {

			res, err := models.CasbinEnforcer.HasNamedPolicy("p", role, "global", path, methods)
			if err != nil {
				return nil, nil, err
			}
			if res {
				existPolicies = append(existPolicies, []string{role, "global", path, methods})
			} else {
				newDiffPolicies = append(newDiffPolicies, []string{role, "global", path, methods})
			}
		}
	}
	oldDiffPolicies = tools.DiffSlices(oldPolicies, existPolicies)
	return oldDiffPolicies, newDiffPolicies, nil
}

// rbac group策略，g,user,role,dom
func (r *RBACPlatform) AddGroupPolicies(user string, roles []string) error {
	res, err := models.CasbinEnforcer.AddGroupingPolicies(r.GetGroupPoliciesForUser(user, roles))
	if err != nil || !res {
		return errors.New(fmt.Sprintf("Add GroupPolicy for user(%s) failed: %v", user, err))
	}
	return nil
}

func (r *RBACPlatform) UpdateGroupPolicies(user string, roles []string) error {
	oldRoles, err := r.GetGroupPolicies(user)
	if err != nil {
		return err
	}
	existRoles := make([][]string, 0)
	for _, role := range roles {
		res, err := models.CasbinEnforcer.HasNamedGroupingPolicy("g", user, role, "global")
		if err != nil {
			return err
		}
		if res {
			existRoles = append(existRoles, []string{user, role, "global"})
		} else {
			res, err := models.CasbinEnforcer.AddGroupingPolicy(user, role, "global")
			if err != nil {
				return err
			}
			if !res {
				return errors.New("Add GroupPolicy failed in UpdateGroupPolicy")
			}
		}
	}
	deletedRoles := tools.DiffSlices(oldRoles, existRoles)
	if len(deletedRoles) > 0 {
		res, err := models.CasbinEnforcer.RemoveGroupingPolicies(deletedRoles)
		if err != nil {
			return err
		}
		if !res {
			return errors.New("Remove GroupPolicy failed in UpdateGroupPolicy")
		}
	}
	return nil
}

func (r *RBACPlatform) GetGroupPolicies(user string) ([][]string, error) {
	return models.CasbinEnforcer.GetFilteredNamedGroupingPolicy("g", 0, user, "global")
}

func (r *RBACPlatform) RemoveGroupPolicies(user string) error {
	res, err := models.CasbinEnforcer.DeleteRolesForUser(user, "global")
	if err != nil {
		return err
	}
	if !res {
		return errors.New(fmt.Sprintf("Roles not fount for user: %s", user))
	}
	return nil
}

func (r *RBACPlatform) GetGroupPoliciesForUser(user string, roles []string) [][]string {
	res := make([][]string, 0)
	for _, role := range roles {
		res = append(res, []string{user, role, "global"})
	}
	return res
}

func (r *RBACPlatform) GetDiffGroupPolicies(user string, roles []string) ([][]string, [][]string, error) {
	oldDiffGroupPolicies, newDiffGroupPolicies := make([][]string, 0), make([][]string, 0)
	existGroupPolicies := make([][]string, 0)
	oldRoles, err := r.GetGroupPolicies(user)
	if err != nil {
		return nil, nil, err
	}
	for _, role := range roles {
		res, err := models.CasbinEnforcer.HasNamedGroupingPolicy("g", user, role, "global")
		if err != nil {
			return nil, nil, err
		}
		if res {
			existGroupPolicies = append(existGroupPolicies, []string{user, role, "global"})
		} else {
			newDiffGroupPolicies = append(newDiffGroupPolicies, []string{user, role, "global"})
		}
	}
	oldDiffGroupPolicies = tools.DiffSlices(oldRoles, existGroupPolicies)
	return oldDiffGroupPolicies, newDiffGroupPolicies, nil
}

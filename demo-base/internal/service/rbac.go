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

type RBAC struct {
	Roles map[string]map[string][]string `json:"roles"`
}

func init() {
	NewRBAC()
}

func NewRBAC() *RBAC {
	input, err := os.ReadFile(conf.RBACPolicy)
	if err != nil {
		logger.Panic(err.Error())
	}
	var rbac *RBAC

	if err = sonic.Unmarshal(input, &rbac.Roles); err != nil {
		logger.Panic(err.Error())
	}
	return rbac
}

// rbac 角色策略，p, sub, obj, act
func (r *RBAC) AddPolicy(role string, permissions []string) error {
	res, err := models.CasbinEnforcer.AddPoliciesEx(r.GetPoliciesForUser(role, permissions))
	if err != nil || !res {
		return errors.New(fmt.Sprintf("Add Policy for role(%s) failed: %v", role, err))
	}
	return nil
}

func (r *RBAC) UpdatePolicies(role string, permissions []string) error {
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

func (r *RBAC) GetPolicies(role string) ([][]string, error) {
	return models.CasbinEnforcer.GetFilteredPolicy(0, role)
}

func (r *RBAC) RemovePolicies(role string) error {
	res, err := models.CasbinEnforcer.DeletePermissionsForUser(role)
	if err != nil {
		return err
	}
	if !res {
		return errors.New(fmt.Sprintf("Policies not fount for role: %s", role))
	}
	return nil
}

func (r *RBAC) AddPoliciesEx(rules [][]string) error {
	res, err := models.CasbinEnforcer.AddPoliciesEx(rules)
	if err != nil || !res {
		return errors.New("AddPolicies failed")
	}
	return nil
}
func (r *RBAC) RemovePoliciesEx(rules [][]string) error {
	res, err := models.CasbinEnforcer.RemovePolicies(rules)
	if err != nil || !res {
		return errors.New("RemovePolicies failed")
	}
	return nil
}

func (r *RBAC) GetPoliciesForUser(user string, permissions []string) [][]string {
	res := make([][]string, 0)
	for _, permission := range permissions {
		for path, methods := range r.Roles[permission] {
			for _, method := range methods {
				res = append(res, []string{user, path, method})
			}
		}
	}
	return res
}

func (r *RBAC) GetDiffPolicies(role string, permissions []string) ([][]string, [][]string, error) {
	oldDiffPolicies, newDiffPolicies := make([][]string, 0), make([][]string, 0)
	existPolicies := make([][]string, 0)
	oldPolicies, err := r.GetPolicies(role)
	if err != nil {
		return nil, nil, err
	}
	for _, permission := range permissions {
		objects := r.Roles[permission]
		for path, methods := range objects {
			for _, method := range methods {
				res, err := models.CasbinEnforcer.HasPolicy(role, path, method)
				if err != nil {
					return nil, nil, err
				}
				if res {
					existPolicies = append(existPolicies, []string{role, path, method})
				} else {
					newDiffPolicies = append(newDiffPolicies, []string{role, path, method})
				}
			}
		}
	}
	oldDiffPolicies = tools.DiffSlices(oldPolicies, existPolicies)
	return oldDiffPolicies, newDiffPolicies, nil
}

// rbac group策略，g,user,role
func (r *RBAC) AddGroupPolicy(user string, roles []string) error {
	res, err := models.CasbinEnforcer.AddGroupingPolicies(r.GetGroupPoliciesForUser(user, roles))
	if err != nil || !res {
		return errors.New(fmt.Sprintf("Add GroupPolicy for user(%s) failed: %v", user, err))
	}
	return nil
}

func (r *RBAC) UpdateGroupPolicies(user string, roles []string) error {
	oldRoles, err := r.GetGroupPolicies(user)
	if err != nil {
		return err
	}
	existRoles := make([][]string, 0)
	for _, role := range roles {
		res, err := models.CasbinEnforcer.HasGroupingPolicy(user, role)
		if err != nil {
			return err
		}
		if res {
			existRoles = append(existRoles, []string{user, role})
		} else {
			res, err := models.CasbinEnforcer.AddGroupingPolicy(user, role)
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

func (r *RBAC) GetGroupPolicies(user string) ([][]string, error) {
	return models.CasbinEnforcer.GetFilteredGroupingPolicy(0, user)
}

func (r *RBAC) RemoveGroupPolicies(user string) error {
	res, err := models.CasbinEnforcer.DeleteRolesForUser(user)
	if err != nil {
		return err
	}
	if !res {
		return errors.New(fmt.Sprintf("Roles not fount for user: %s", user))
	}
	return nil
}

func (r *RBAC) GetGroupPoliciesForUser(user string, roles []string) [][]string {
	res := make([][]string, 0)
	for _, role := range roles {
		res = append(res, []string{user, role})
	}
	return res
}

func (r *RBAC) GetDiffGroupPolicies(user string, roles []string) ([][]string, [][]string, error) {
	oldDiffGroupPolicies, newDiffGroupPolicies := make([][]string, 0), make([][]string, 0)
	existGroupPolicies := make([][]string, 0)
	oldRoles, err := r.GetGroupPolicies(user)
	if err != nil {
		return nil, nil, err
	}
	for _, role := range roles {
		res, err := models.CasbinEnforcer.HasGroupingPolicy(user, role)
		if err != nil {
			return nil, nil, err
		}
		if res {
			existGroupPolicies = append(existGroupPolicies, []string{user, role})
		} else {
			newDiffGroupPolicies = append(newDiffGroupPolicies, []string{user, role})
		}
	}
	oldDiffGroupPolicies = tools.DiffSlices(oldRoles, existGroupPolicies)
	return oldDiffGroupPolicies, newDiffGroupPolicies, nil
}

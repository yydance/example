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

func (r *RBAC) AddPolicy(role string, permissions []string) error {
	errs := make([]error, 0)
	for _, permission := range permissions {
		objects := r.Roles[permission]
		for path, methods := range objects {
			for _, method := range methods {
				result, err := models.CasbinEnforcer.AddPolicy(role, path, method)
				if err != nil {
					errs = append(errs, err)
				}
				if !result {
					errs = append(errs, errors.New(fmt.Sprintf("Policy exists: %s %s %s", role, path, method)))
				}
			}
		}
	}
	if len(errs) > 0 {
		return errors.New(fmt.Sprintf("Add Policy failed: %v", errs))
	}
	return nil
}

func (r *RBAC) UpdatePolicies(role string, permissions []string) error {
	oldPolicies, err := r.GetPolicies(role)
	if err != nil {
		return err
	}
	existPolicies := make([][]string, 0)
	for _, permission := range permissions {
		objects := r.Roles[permission]
		for path, methods := range objects {
			for _, method := range methods {
				res, err := models.CasbinEnforcer.HasPolicy(role, path, method)
				if err != nil {
					return err
				}
				if res {
					existPolicies = append(existPolicies, []string{role, path, method})
				} else {
					res, err := models.CasbinEnforcer.AddPolicy(role, path, method)
					if err != nil {
						return err
					}
					if !res {
						return errors.New("Add Policy failed in UpdatePolicy")
					}
				}
			}
		}
	}
	deletedPolicies := tools.DiffSlices(oldPolicies, existPolicies)
	if len(deletedPolicies) > 0 {
		res, err := models.CasbinEnforcer.RemovePolicies(deletedPolicies)
		if err != nil {
			return err
		}
		if !res {
			return errors.New("Remove Policy failed in UpdatePolicy")
		}
	}
	return nil
}

func (r *RBAC) GetPolicies(role string) ([][]string, error) {
	return models.CasbinEnforcer.GetFilteredPolicy(0, role)
}

func (r *RBAC) DeletePolicies(role string) error {
	res, err := models.CasbinEnforcer.DeletePermissionsForUser(role)
	if err != nil {
		return err
	}
	if !res {
		return errors.New(fmt.Sprintf("Policies not fount for role: %s", role))
	}
	return nil
}

func (r *RBAC) AddGroupPolicy(user string, roles []string) error {
	errs := make([]error, 0)
	for _, role := range roles {
		result, err := models.CasbinEnforcer.AddGroupingPolicy(user, role)
		if err != nil {
			errs = append(errs, err)
			return err
		}
		if !result {
			errs = append(errs, errors.New(fmt.Sprintf("Policy exists: %s %s", user, role)))
		}
	}
	if len(errs) > 0 {
		return errors.New(fmt.Sprintf("Add GroupPolicy failed: %v", errs))
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

func (r *RBAC) DeleteGroupPolicies(user string) error {
	res, err := models.CasbinEnforcer.DeleteRolesForUser(user)
	if err != nil {
		return err
	}
	if !res {
		return errors.New(fmt.Sprintf("Roles not fount for user: %s", user))
	}
	return nil
}

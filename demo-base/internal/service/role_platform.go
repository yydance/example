package service

import (
	"demo-base/internal/models"
	"errors"
	"fmt"
)

type RolePlatform struct {
	Name        string   `json:"name" validate:"required,max=64"`
	Description string   `json:"description" validate:"max=256"`
	Permissions []string `json:"permissions"`
}

func (r *RolePlatform) Create() error {
	err := validate.Struct(r)
	if err != nil {
		return errors.New("invalid input")
	}
	var rbac *RBAC
	if err := rbac.AddPolicy(r.Name, r.Permissions); err != nil {
		return err
	}

	role := models.RolePlatform{
		Name:        r.Name,
		Description: r.Description,
		Permissions: r.Permissions,
	}
	if err := role.Create(); err != nil {
		if err := rbac.RemovePolicies(r.Name); err != nil {
			return err
		} // rollback
		return err
	}

	return nil
}

func (r *RolePlatform) Update() error {
	err := validate.Struct(r)
	if err != nil {
		return errors.New("invalid input")
	}

	var rbac *RBAC
	oldDiffPolicies, newDiffPolicies, err := rbac.GetDiffPolicies(r.Name, r.Permissions)
	if err != nil {
		return err
	}
	if err := rbac.UpdatePolicies(r.Name, r.Permissions); err != nil {
		return err
	}
	role := models.RolePlatform{
		Name:        r.Name,
		Description: r.Description,
		Permissions: r.Permissions,
	}
	// rollback，更新失败，添加删除的策略，删除新增的策略

	if err := role.Update(); err != nil {
		if len(oldDiffPolicies) > 0 {
			if err := rbac.AddPoliciesEx(oldDiffPolicies); err != nil {
				return err
			}
		}
		if len(newDiffPolicies) > 0 {
			if err := rbac.RemovePoliciesEx(newDiffPolicies); err != nil {
				return err
			}
		}
	}
	return nil
}

func (r *RolePlatform) Delete() error {
	var rbac *RBAC
	policies, err := rbac.GetPolicies(r.Name)
	if err != nil {
		return errors.New((fmt.Sprintf("GetPolicies for role(%s) error: %v", r.Name, err)))
	}
	if err := rbac.RemovePolicies(r.Name); err != nil {
		return err
	}
	role := models.RolePlatform{
		Name: r.Name,
	}
	// rollback，删除失败，需要恢复rbac中的数据，先查询出原来的数据，再进行恢复

	if err := role.Delete(); err != nil {
		res, err := models.CasbinEnforcer.AddPoliciesEx(policies)
		if err != nil {
			return err
		}
		if !res {
			return errors.New("AddPoliciesEx failed")
		}
		return err
	}
	return nil
}

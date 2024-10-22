package models

import (
	"demo-base/internal/conf"
	"demo-base/internal/utils/logger"

	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type RBACRule struct {
	gorm.Model
	Ptype string `gorm:"max=2"`
	V0    string `gorm:"max=64"`
	V1    string `gorm:"max=64"`
	V2    string `gorm:"max=8"`
}

/*
func (RBACRule) TableName() string {
	return "rbac_rules"
}
*/

var CasbinEnforcer *casbin.Enforcer

func InitCasbinEnforcer() {
	gadapter, err := gormadapter.NewAdapterByDBWithCustomTable(DB, &RBACRule{}, "rbac_rules")
	if err != nil {
		logger.Error("rbac init error: %v", err)
	}
	CasbinEnforcer, err = casbin.NewEnforcer(conf.RBACModel, gadapter)
	if err != nil {
		logger.Error("casbin new enforcer error: %v", err)
	}
}

/*
type RBAC struct{}

func (r *RBAC) Enforce(user, resource, action string) (bool, error) {
	return CasbinEnforcer.Enforce(user, resource, action)
}

func (r *RBAC) AddPolicy(rule []string) (bool, error) {
	return CasbinEnforcer.AddPolicy(rule)
}

func (r *RBAC) AddPolicies(rules [][]string) (bool, error) {
	return CasbinEnforcer.AddPolicies(rules)
}

func (r *RBAC) RemovePolicy(rule []string) (bool, error) {
	return CasbinEnforcer.RemovePolicy(rule)
}

func (r *RBAC) RemovePolicies(rules [][]string) (bool, error) {
	return CasbinEnforcer.RemovePolicies(rules)
}
*/

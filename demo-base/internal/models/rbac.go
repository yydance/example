package models

import (
	"demo-base/internal/utils/logger"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"gorm.io/gorm"
)

type RBACRule struct {
	gorm.Model
	Ptype string `gorm:"max=4"`
	V0    string `gorm:"max=64"`
	V1    string `gorm:"max=64"`
	V2    string `gorm:"max=64"`
	V3    string `gorm:"max=64"`
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
	m, err := model.NewModelFromString(`
	[request_definition]
	r = sub, obj, act

	[policy_definition]
	p = sub, obj, act

	[role_definition]
	g = _, _

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = r.sub == p.sub && g(r.sub,p.sub) && r.obj == p.obj && r.act == p.act || r.sub == "root" || r.sub == "admin" || r.sub == "Admin"
	`)
	if err != nil {
		logger.Error("casbin new model error: %v", err)
	}
	CasbinEnforcer, err = casbin.NewEnforcer(m, gadapter)
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

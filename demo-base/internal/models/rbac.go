package models

import (
	"demo-base/internal/utils/logger"

	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/util"
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

func (RBACRule) TableName() string {
	return "rbac_rules"
}

var CasbinEnforcer *casbin.Enforcer

func InitCasbinEnforcer() {
	gadapter, err := gormadapter.NewAdapterByDBWithCustomTable(DB, &RBACRule{}, "rbac_rules")
	if err != nil {
		logger.Error("rbac init error: %v", err)
	}
	m, err := model.NewModelFromString(`
	[request_definition]
	r = sub, dom, obj, act

	[policy_definition]
	p = sub, dom, obj, act

	[role_definition]
	g = _, _, _

	[policy_effect]
	e = some(where (p.eft == allow))

	[matchers]
	m = r.dom == p.dom && g(r.sub,p.sub,r.dom) && r.obj == p.obj && regexMatch(r.act,p.act)
	`)
	if err != nil {
		logger.Error("casbin new model error: %v", err)
	}
	CasbinEnforcer, err = casbin.NewEnforcer(m, gadapter)
	if err != nil {
		logger.Error("casbin new enforcer error: %v", err)
	}
	CasbinEnforcer.AddNamedMatchingFunc("p", "KeyMatch2", util.KeyMatch2)
	//CasbinEnforcer.AddNamedDomainMatchingFunc("g", "globMatch", globMatch)
}

func globMatch(key1 string, key2 string) bool {
	if ok, err := util.GlobMatch(key1, key2); err == nil {
		return ok
	}
	return false
}

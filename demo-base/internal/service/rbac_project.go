package service

import (
	"demo-base/internal/conf"
	"demo-base/internal/utils/logger"
	"os"

	"github.com/bytedance/sonic"
)

type RBACProject struct {
	Roles map[string]map[string][]string `json:"roles"`
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

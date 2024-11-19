package opa

import (
	"demo-base/internal/conf"
	"demo-base/internal/models"
	"fmt"

	"github.com/bytedance/sonic"
	mapset "github.com/deckarep/golang-set"
)

func genPolicy(username string) error {
	var user_roles = map[string]any{}

	if err := sonic.Unmarshal([]byte(role_grants), &policies); err != nil {
		return nil
	}

	key := conf.RolesPrefix + "/" + username
	roles, err := models.CacheStore.ValueToMap(key)
	if err != nil {
		//policies["roles"] = make(map[string]interface{})
		return fmt.Errorf("get user roles failed: %v", err)
	}

	user_roles[username] = roles
	policies["roles"] = user_roles

	return nil
}

func isPlatformView(username string) bool {
	user_roles, ok := policies["roles"].(map[string]interface{})
	if !ok {
		return false
	}
	roles, ok := user_roles[username].([]interface{})
	if !ok {
		return false
	}

	s := mapset.NewSetFromSlice(roles)

	return s.Contains("Platform:View")
}

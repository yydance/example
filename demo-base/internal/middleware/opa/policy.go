package opa

import (
	"demo-base/internal/conf"
	"demo-base/internal/models"

	"github.com/bytedance/sonic"
	mapset "github.com/deckarep/golang-set"
)

func genPolicy(username string) error {
	var user_roles = map[string]any{}
	if err := sonic.Unmarshal([]byte(role_grants), &policies); err != nil {
		return nil
	}

	key := conf.RolesPrefix + "/" + username
	value, ok := models.CacheStore.Get(key)
	if !ok {
		policies["user_roles"] = make(map[string]interface{})
		return nil
	}
	user_roles[username] = value
	policies["user_roles"] = user_roles

	return nil
}

func isPlatformView(username string) bool {
	user_roles, ok := policies["user_roles"].(map[string]interface{})
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

package opa

import (
	"github.com/bytedance/sonic"
	mapset "github.com/deckarep/golang-set"
)

func genPolicy(username string) error {
	var user_roles = map[string]any{}
	var value = map[string]any{}
	if err := sonic.Unmarshal([]byte(role_grants), &policies); err != nil {
		return nil
	}

	value["global"] = []string{"Platform:Admin"}

	user_roles[username] = value
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

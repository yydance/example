package opa

var (
	policies = map[string]interface{}{}
	module   = `
		package authz

		import rego.v1

		default allow := false

		allow if "Platform:Admin" in user_roles

		allow if {
			input.domain != "global"
			some grant in user_is_granted
			input.action in grant.action
			input.object == grant.object
		}

		allow if {
			input.domain != "global"
			some grant in user_is_granted
			input.action in grant.action
			regex.globs_match(grant.object, input.object)
		}

		allow if {
			input.domain == "global"
			some grant in user_is_granted
			input.action in grant.action
			input.object == grant.object
		}

		allow if {
			input.domain == "global"
			some grant in user_is_granted
			input.action in grant.action
			regex.globs_match(grant.object, input.object)
		}

		user_roles := role if {
			some kv 
			kv = data.roles[input.username]
			some role 
			role = kv[input.domain]
		}

		user_is_granted contains grant if {
			some role in user_roles
			some grant
			grant = data.role_grants[role][_]
		}
	`

	role_grants = `{
		"role_grants": {
			"Platform:View": [{"action": ["GET"], "object": "/*"}],
			"Platform:Admin": [{"action": "*", "object":"*"}],
			"Platform:Users:View": [
				{"action": ["GET"], "object": "/api/v1/users"},
				{"action": ["GET"], "object": "/api/v1/users/[0-9a-z]+"}
			],
			"Platform:Users:Management": [
				{"action": ["GET","POST"], "object": "/api/v1/users"},
				{"action": ["GET","PUT","DELETE"], "object": "/api/v1/users/[0-9a-z+]+"}
			],
			"Platform:Roles:View": [
				{"action": ["GET"], "object": "/api/v1/roles"},
				{"action": ["GET"], "object": "/api/v1/roles/[0-9a-z]+"}
			],
			"Platform:Roles:Management": [
				{"action": ["GET","POST"], "object": "/api/v1/roles"},
				{"action": ["GET","PUT","DELETE"], "object": "/api/v1/roles/[0-9a-z]+"}
			],
			"Platform:Projects:View": [
				{"action": ["GET"], "object": "/api/v1/projects"},
				{"action": ["GET"], "object": "/api/v1/projects/[0-9a-z]+"}
			],
			"Platform:Projects:Management": [
				{"action": ["GET","POST"], "object": "/api/v1/platform/projects"},
				{"action": ["GET","PUT","DELETE"], "object": "/api/v1/platform/projects/[0-9a-z]+"},
				{"action": ["GET","POST"], "object": "/api/v1/projects/[0-9a-z]+/roles"},
				{"action": ["GET","PUT","DELETE"], "object": "/api/v1/projects/[0-9a-z]+/roles/[0-9a-z]+"},
				{"action": ["GET","POST"], "object": "/api/v1/projects/[0-9a-z]+/members"},
				{"action": ["GET","PUT","DELETE"], "object": "/api/v1/projects/[0-9a-z]+/members/[0-9a-z]+"},
				{"action": ["GET"], "object": "/api/v1/projects/[0-9a-z]+/summary"},
				{"action": ["GET"], "object": "/api/v1/apps"},
				{"action": ["GET"], "object": "/api/v1/apps/[0-9a-z]+"}
			],
			"Proj:Roles:View": [
				{"action": ["GET"], "object": "/api/v1/projects/[0-9a-z]+/roles"},
				{"action": ["GET"], "object": "/api/v1/projects/[0-9a-z]+/members"}
			],
			"Proj:Roles:Management": [
				{"action": ["GET"], "object": "/api/v1/users"},
				{"action": ["GET","POST"], "object": "/api/v1/projects/[0-9a-z]+/roles"},
				{"action": ["GET","PUT","DELETE"], "object": "/api/v1/projects/[0-9a-z]+/roles/[0-9a-z]+"},
				{"action": ["GET","POST"], "object": "/api/v1/projects/[0-9a-z]+/members"},
				{"action": ["GET","PUT","DELETE"], "object": "/api/v1/projects/[0-9a-z]+/members/[0-9a-z]+"}
			],
			"Proj:Project:Management": [
				{"action": ["GET","PUT"], "object": "/api/v1/projects/[0-9a-z]+"},
				{"action": ["GET","POST"], "object": "/api/v1/projects/[0-9a-z]+/roles"},
				{"action": ["GET","PUT","DELETE"], "object": "/api/v1/projects/[0-9a-z]+/roles/[0-9a-z]+"},
				{"action": ["GET","POST"], "object": "/api/v1/projects/[0-9a-z]+/members"},
				{"action": ["GET","PUT","DELETE"], "object": "/api/v1/projects/[0-9a-z]+/members/[0-9a-z]+"},
				{"action": ["GET"], "object": "/api/v1/projects/[0-9a-z]+/summary"},
				{"action": ["GET"], "object": "/api/v1/apps"},
				{"action": ["GET"], "object": "/api/v1/apps/[0-9a-z]+"}
			],
			"Proj:Project:View": [
				{"action": ["GET"], "object": "/api/v1/projects/[0-9a-z]+"},
				{"action": ["GET"], "object": "/api/v1/projects/[0-9a-z]+/summary"},
				{"action": ["GET"], "object": "/api/v1/apps"},
				{"action": ["GET"], "object": "/api/v1/apps/[0-9a-z]+"},
				{"action": ["GET"], "object": "/api/v1/projects/[0-9a-z]+/roles"},
				{"action": ["GET"], "object": "/api/v1/projects/[0-9a-z]+/members"}
			]
		}
	}`
)

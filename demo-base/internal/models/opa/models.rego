package rbac

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





---

### etcd策略存储结构
```
/demo_base, init_dir
  /user_roles, init_dir
    /username, {"global": ["Platform:Users:View","Platform:Projects:View"], "qa": ["Proj:Roles:Management"]}
  /role_grants, init_dir
    /rolename, [{"action": "*", "object":"*"}]
```

example:
```
/demo_base, init
/demo_base/user_roles, init
/demo_base/user_roles/damon, {"global": ["Platform:Users:View","Platform:Projects:View"], "qa": ["Proj:Roles:Management"]}
/demo_base/user_roles/bob, {"global": ["Platform:Users:View","Platform:Projects:View"], "qa": ["Proj:Roles:Management"]}

/demo_base/role_grants, init
/demo_base/role_grants/Platform:Admin, [{"action": "*", "object":"*"}]
/demo_base/role_grants/Platform:Users:View, [{"action": "*", "object":"*"}]
```
配置文件描述及权限说明

> 该文档camin rbac模型已废弃，请参考models内opa rbac模型

### 配置文件
待补充...

### 权限
权限模型为RBAC，即角色-权限-用户模型，分为全局角色和项目角色，有其一即可访问。

#### 模型
rbac_demo.conf
```
[request_definition]
r = sub, obj, act

[policy_definition]
p = sub, obj, act

[role_definition]
g = _, _

[policy_effect]
e = some(where (p.eft == allow)) && !some(where (p.eft == deny))

[matchers]
m = r.sub == p.sub && g(r.sub,p.sub) && r.obj == p.obj && r.act == p.act || r.sub == "root" || r.sub == "admin"
```

#### 策略
示例
```
p, user1, /api/v1/users/*, write
p, user2, /api/v1/users/*, read
p, admin_groups, /*, admin
p, view_groups, /*, read
p, project_admin, /api/v1/projects/:name, admin
p, user1, /api/v1/projects/:name/users, write
p, viewer1, /api/v1/projects/:name, read

g, damon, admin_groups
g, viewer, view_groups
```

#### 权限详情

| 权限 | 说明 |
| :--- | :--- |
| admin | 管理员权限，拥有所有权限 |
| read | 读权限，可以查看资源 |
| write | 写权限，可以修改资源 |
| delete | 删除权限，可以删除资源 |

```go
var permissions = map[string][]string{
	"admin": []string{
	    "GET",
	    "POST",
	    "PUT",
	    "DELETE",
        "PATCH",
        "CREATE",
	},
    "read": []string{
        "GET",
    },
    "write": []string{
        "GET",
        "POST",
        "PUT",
        "PATCH",
        "CREATE",
    },
    "delete": []string{"DELETE"}
}

var roles_platform = map[string]any{
	"User Manager": {
		"/api/v1/users": ["GET", "POST"],
		"/api/v1/users/:name": ["GET","PUT","PATCH","DELETE"],
		"/api/v1/users/:name/updatePassword": ["PUT"],
	},
	"Roles Manager": {
		"/api/v1/roles": ["GET", "POST"],
		"/api/v1/roles/:name": ["GET","PUT","PATCH","DELETE"],
	},
	"Projects Manager": {
		"/api/v1/projects": ["GET","POST"],
		"/api/v1/projects/:name": ["GET","PUT","PATCH","DELETE"],
		"/api/v1/projects/:name/roles": ["GET","POST"],
		"/api/v1/projects/:name/roles/:roleName": ["GET","PUT","PATCH","DELETE"],
		"/api/v1/projects/:name/users": ["GET","POST"],
		"/api/v1/projects/:name/users/:userName": ["GET","PUT","PATCH","DELETE"],
	}
}

var roles_project = map[string][]string{
    "Project Manager": []string{
        "/api/v1/projects/:name",
    }
}
```

### 路由列表
```
- /api/v1  
  - /login  
  - /logout  
  - /users, GET,POST  
    - /:name，GET,PUT,DELETE,PATCH  
    - /search, GET  
  - /projects, GET,POST  
    - /:name, GET,PUT,DELETE,PATCH 
      - /summary, GET 
      - /users, GET,POST  
        - /:name, GET,PUT,DELETE,PATCH
      - roles,GET,POST
        - /:name, GET,PUT,DELETE,PATCH
      - /apps
  - /settings, GET,POST

```
  
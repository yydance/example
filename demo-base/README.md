demo-base
---
基于gofiber框架

### 项目结构


### 运行

### Notice
应用启动后会初始化表结构，并初始化以下数据：
- 超级用户admin，密码abc123456，绑定了平台角色admin
- 平台角色admin，权限为所有权限
- 默认项目default，其管理员为admin，项目角色`Project View`，`Project Admin`，`Project Developer`
  - project view角色，权限为查看项目
  - project admin角色，权限为管理项目
  - project developer角色，权限为开发项目

### 权限模型
权限分为`平台权限`和`项目权限`。  
平台权限包括：
- 用户管理，`User Management`->`/api/v1/users/*`
- 角色管理，`Platform Role Management`->`/api/v1/roles/*`
- 项目管理，`Project Management`->`/api/v1/projects/*`
- 系统配置，`System Config`->`/api/v1/system/*`
- ...  

项目权限包括：
- 项目查看，`Project View`->`/api/v1/projects/:name/*`
- 应用管理，`App Management`->`/api/v1/projects/:name/apps/*`
- 角色管理，`Role Management`->`/api/v1/projects/:name/roles/*`，注意：该角色与平台角色不同，该角色仅用于项目内部角色管理，同时该角色管理员可以添加成员并为成员分配角色
- ...

### TODO

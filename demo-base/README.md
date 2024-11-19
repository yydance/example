demo-base
---
基于gofiber框架

### 项目结构
TODO

### 运行

### Notice
应用启动后会初始化表结构，并初始化以下数据：
- 超级用户admin，密码abc123456，绑定了平台角色Platform:Admin
- 平台角色Platform:Admin，权限为所有权限
- 默认项目default，其管理员为admin，项目角色`Proj:Project:View`，`Proj:Project:Management`，`Proj:Roles:Management`，项目中的成员自动继承项目View角色
  

### 权限模型
权限分为`平台角色`和`项目角色`。  
#### 平台角色，面向平台管理者
- 用户管理，`Platform:Users:Management`->`/api/v1/users/*`->`GET,POST,PUT,DELETE`
- 角色管理，`Platform:Roles:Management`->`/api/v1/roles/*`->`GET,POST,PUT,DELETE`
- 项目管理，`Platform:Project:Management`->`/api/v1/projects/*`->`GET,POST,PUT,DELETE`
- 系统配置，`Platform:System`->`/api/v1/system/*`
- 项目浏览，`Platform:Project:View`->`/api/v1/projects` ->`GET`
- ...

#### 项目角色，面向项目管理使用者
- 项目查看，`Proj:Project:View`->`/api/v1/projects*`->`GET`，项目成员自动继承
- 应用管理，`Proj:Apps:Management`->`/api/v1/apps/*`->`GET,POST,PUT,DELETE`
- 角色管理，`Proj:Roles:Management`->`[{"/api/v1/projects/:name/roles/*":"GET,POST,PUT,DELETE"},{"/api/v1/users":"GET"}]`，注意：该角色与平台角色不同，该角色仅用于项目内部角色管理，同时该角色管理员可以添加成员并为成员分配角色
- ...

#### 角色优先级
- 平台角色 > 项目角色
- 平台只读角色例外，其优先级最低
- 实现，权限校验时，先检查平台角色，再检查项目角色，最后检查只读角色

### TODO
- [x] 项目管理

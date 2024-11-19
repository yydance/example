package service

import (
	"context"
	"demo-base/internal/conf"
	"demo-base/internal/models"
	"demo-base/internal/utils/logger"
	"errors"

	"github.com/bytedance/sonic"
)

// Service is the interface for the service layer
// It should be implemented by any service that needs to interact with the database

type UserInput struct {
	Name     string   `json:"name" validate:"required,max=64"`
	Email    string   `json:"email" validate:"omitempty,email"`
	Roles    []string `json:"roles" validate:"omitempty"`
	Password string   `json:"password" validate:"omitempty,min=6,max=64,alphanum"`
}

type UserPassword struct {
	Name     string `json:"name" validate:"required"`
	Password string `json:"password" validate:"required,min=6,max=64,alphanum"`
}

type UsersOutput struct {
	Name      string   `json:"name"`
	Email     string   `json:"email"`
	Roles     []string `json:"roles"`
	CreatedAt string   `json:"created_at"`
	UpdatedAt string   `json:"updated_at"`
}

func (u *UserInput) Create() error {
	err := validate.Struct(u)
	if err != nil || u.Password == "" {
		return errors.New("invalid input")
	}

	user := models.User{
		Name:     u.Name,
		Email:    u.Email,
		Password: u.Password,
		Roles:    u.Roles,
	}

	if err := user.Create(); err != nil {
		return err
	}

	permissions, err := u.GetAllPermissions()
	if err != nil {
		return errors.New("failed to get permissions")
	}
	// 将用户名u.Name、预定义角色写入etcd global策略
	// key: /prefix/roles/{u.name}/
	// value: {"global": roles}
	key := conf.RolesPrefix + "/" + u.Name
	var value map[string][]string

	getValue, ok := models.CacheStore.Get(key)
	if ok {
		err = sonic.Unmarshal([]byte(getValue.(string)), &value)
		if err != nil {
			return err
		}
		value["global"] = permissions
	} else {
		value = map[string][]string{"global": permissions}
	}

	roles, err := sonic.Marshal(value)
	if err != nil {
		return err
	}

	if err := models.EtcdStorage.Set(context.TODO(), key, roles, 0); err != nil {
		return err
	}

	return nil
}

func (u *UserInput) Get() (models.User, error) {
	user := models.User{
		Name: u.Name,
	}
	return user, user.Find()
}

func (u *UserInput) Update() error {
	err := validate.Struct(u)
	if err != nil {
		return errors.New("invalid input")
	}
	user := models.User{
		Name:  u.Name,
		Email: u.Email,
		Roles: u.Roles,
	}
	if err := user.Update(); err != nil {
		return err
	}
	permissions, err := u.GetAllPermissions()
	if err != nil {
		return errors.New("failed to get permissions")
	}
	key := conf.RolesPrefix + "/" + u.Name
	var value map[string][]string

	getValue, ok := models.CacheStore.Get(key)
	if ok {
		err = sonic.Unmarshal([]byte(getValue.(string)), &value)
		if err != nil {
			return err
		}
		value["global"] = permissions
	} else {
		return errors.New("data not found")
	}
	roles, err := sonic.Marshal(value)
	if err != nil {
		return err
	}
	if err := models.EtcdStorage.Set(context.TODO(), key, roles, 0); err != nil {
		return err
	}
	return nil
}

func (u *UserPassword) UpdatePassword() error {
	user := models.User{
		Name:     u.Name,
		Password: u.Password,
	}
	return user.UpdatePassword()
}

func (u *UserInput) Delete() error {
	user := models.User{
		Name: u.Name,
	}
	if err := user.Delete(); err == nil {
		key := conf.RolesPrefix + "/" + u.Name
		var value map[string][]string
		getValue, err := models.EtcdStorage.Get(context.TODO(), key)
		if err != nil {
			return err
		}
		err = sonic.Unmarshal(getValue, &value)
		if err != nil {
			return err
		}
		delete(value, "global")
		if len(value) == 0 {
			if err := models.EtcdStorage.Delete(context.TODO(), key); err != nil {
				return err
			}
			return nil
		}
		roles, err := sonic.Marshal(value)
		if err != nil {
			return err
		}
		if err := models.EtcdStorage.Set(context.TODO(), key, roles, 0); err != nil {
			return err
		}
		return nil
	}

	return nil
}

func (u *UserInput) List(pageNum, pageSize int) ([]UsersOutput, error) {
	var user models.User
	users := make([]UsersOutput, 0)
	userList, err := user.List(pageNum, pageSize)
	if err != nil {
		return users, err
	}
	for _, user := range userList {
		users = append(users, UsersOutput{
			Name:      user.Name,
			Email:     user.Email,
			Roles:     user.Roles,
			CreatedAt: user.CreatedAt.Format("2006-01-02 15:04:05"),
			UpdatedAt: user.UpdatedAt.Format("2006-01-02 15:04:05"),
		})
	}
	return users, nil
}

func (u *UserInput) GetAllPermissions() ([]string, error) {
	var r = RolePlatform{}
	permissions := []string{}
	for i := range u.Roles {
		r.Name = u.Roles[i]
		roles, err := r.GetRoles()
		if err != nil {
			return nil, err
		}
		permissions = append(permissions, roles...)
	}

	return permissions, nil
}

func (u *UserInput) IsExist() (models.User, bool) {
	var user = models.User{Name: u.Name}
	return user.IsExist()
}

func InitAdmin() {
	// create admin role
	logger.Info("init admin role")
	role := RolePlatform{
		Name: "Admin",
		Permissions: []string{
			"Platform:Admin",
		},
		Description: "Admin role",
	}
	if err := role.Create(); err != nil {
		logger.Warnf("failed to create admin role: %s", err.Error())
	} else {
		logger.Info("init admin role success")
	}
	// create admin user
	admin := UserInput{
		Name:     "admin",
		Password: "admin0416",
		Roles:    []string{"Admin"},
		Email:    "admin@platform.com",
	}
	if err := admin.Create(); err != nil {
		//return fmt.Errorf("failed to create admin user: %s", err.Error())
		logger.Warnf("failed to create admin user: %s", err.Error())
	} else {
		logger.Info("init admin user success")
	}
}

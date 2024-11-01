package service

import (
	"demo-base/internal/models"
	"demo-base/internal/utils/jwt"
	"errors"
)

type LoginInput struct {
	Username string `json:"username" validate:"required,max=64"`
	Password string `json:"password" validate:"required,max=64"`
}

func (l *LoginInput) Login() (string, error) {
	if err := validate.Struct(l); err != nil {
		return "", err
	}
	u := models.User{}
	user, ok := u.IsExist(l.Username)
	if !ok {
		return "", errors.New("用户不存在")
	}
	if user.Password != l.Password {
		return "", errors.New("密码错误")
	}
	// 生成token
	var jwt *jwt.JWT
	token, err := jwt.GenerateToken(l.Username)
	if err != nil {
		return "", errors.New("生成token失败")
	}
	return token, nil
}

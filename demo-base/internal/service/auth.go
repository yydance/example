package service

import (
	"demo-base/internal/models"
	"demo-base/internal/utils/jwt"
	"demo-base/internal/utils/logger"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type LoginInput struct {
	Username string `json:"username" validate:"required,max=64"`
	Password string `json:"password" validate:"required,max=64"`
}

func (l *LoginInput) Login() (string, error) {
	if err := validate.Struct(l); err != nil {
		return "", err
	}
	u := models.User{Name: l.Username}
	user, ok := u.IsExist()
	if !ok {
		return "", errors.New("用户不存在")
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(l.Password))
	if err != nil {
		return "", errors.New("密码错误")
	}
	// 生成签名
	signature := jwt.NewJWT(l.Username)

	logger.Infof("user %s login success, token: %s", l.Username, signature)
	return signature, nil
}

package user

import (
	"demo-dashboard/internal/core/models"
	"demo-dashboard/internal/handler"
	"demo-dashboard/internal/utils"
	"demo-dashboard/internal/utils/gvalidator"
	jwtToken "demo-dashboard/internal/utils/jwt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type Page struct {
	PageSize int `json:"page_size" validate:"required,page_size" default:"10"`
	PageNum  int `json:"page_num" validate:"required,page_num,min=1" default:"1"`
}

type LoginInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func SignIn(c *fiber.Ctx) error {

	appFiber := handler.Fiber{C: c}
	xvalidator := gvalidator.New()
	login := LoginInput{
		Username: c.FormValue("username"),
		Password: c.FormValue("password"),
	}
	errs := xvalidator.ErrMsgs(login)
	if len(errs) > 0 {
		return appFiber.Handler(fiber.StatusBadRequest, fiber.StatusBadRequest, strings.Join(errs, " and "), nil)
	}

	user := models.User{
		Username: login.Username,
		Password: login.Password,
	}
	userinfo, err := user.Info()
	if err != nil {
		return appFiber.Handler(fiber.StatusUnauthorized, fiber.StatusUnauthorized, "用户名为空或者用户名错误", nil)
	}
	uservalue := userinfo.(models.User)
	if !utils.DecodeHash(uservalue.Password, user.Password) {
		return appFiber.Handler(fiber.StatusUnauthorized, fiber.StatusUnauthorized, "密码错误", nil)
	}

	signToken, err := jwtToken.GenToken(login.Username)

	return appFiber.Handler(fiber.StatusOK, fiber.StatusOK, signToken, err)
}

func SignOut(c *fiber.Ctx) error {

	return nil
}

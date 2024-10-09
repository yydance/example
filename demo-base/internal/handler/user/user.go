package user

import (
	"demo-base/internal/models"

	"github.com/gofiber/fiber/v2"
)

type Page struct {
	pageNum   int `json:"page_num" validate:"required,page_num" default:"10"`
	pageLimit int `json:"page_limit" validate:"required,page_limit" default:"1"`
}

func Create(c *fiber.Ctx) error {
	// TODO
	user := &models.User{}
	if err := c.BodyParser(user); err != nil {
		c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":  err.Error(),
			"code": fiber.StatusBadRequest, // should be customizable code
			"data": nil,
		})
	}

	if err := user.Create(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":  "internal server error",
			"code": fiber.StatusInternalServerError,
			"data": nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "success",
		"code": fiber.StatusOK,
		"data": nil,
	})
}

func Get(c *fiber.Ctx) error {
	//user := &models.User{}
	return nil
}

func Update(c *fiber.Ctx) error {
	var user *models.User
	if err := c.BodyParser(user); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"msg":  err.Error(),
			"code": fiber.StatusBadRequest,
			"data": nil,
		})
	}
	isExist, _ := user.IsExist(user.Name)
	if isExist {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"code": fiber.StatusBadRequest,
			"msg":  "user already exists, please use another name",
			"data": nil,
		})
	}
	if err := user.Update(); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":  "internal server error",
			"code": fiber.StatusInternalServerError,
			"data": nil,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "success",
		"code": fiber.StatusOK,
		"data": nil,
	})
}

func GetAll(c *fiber.Ctx) error {
	page := Page{
		pageNum:   c.QueryInt("page_num"),
		pageLimit: c.QueryInt("page_limit"),
	}

	var user *models.User
	users, err := user.FindByPage(page.pageNum, page.pageLimit)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"msg":  "internal server error",
			"code": fiber.StatusInternalServerError,
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"msg":  "success",
		"data": users,
		"code": fiber.StatusOK,
	})
}

func Delete(c *fiber.Ctx) error {
	return nil
}

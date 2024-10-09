package handler

import (
	"github.com/gofiber/fiber/v2"
)

// NOTICE: Handler is a wrapper for fiber.Ctx, and not be used
type Fiber struct {
	C *fiber.Ctx
}

func (f *Fiber) Handler(httpCode, errCode int, msg string, data interface{}) error {
	return f.C.Status(httpCode).JSON(fiber.Map{
		"code": errCode,
		"msg":  msg,
		"data": data,
	})
}

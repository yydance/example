package handler

import "github.com/gofiber/fiber/v2"

type Fiber struct {
	C *fiber.Ctx
}

func (f *Fiber) Handler(httpCode, errCode int, message string, data any) error {

	return f.C.Status(httpCode).JSON(fiber.Map{
		"code":    errCode,
		"message": message,
		"data":    data,
	})
}

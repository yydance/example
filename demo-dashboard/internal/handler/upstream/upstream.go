package upstream

import (
	"demo-dashboard/internal/conf"
	"demo-dashboard/internal/core/models"
	"demo-dashboard/internal/handler"
	"demo-dashboard/internal/handler/entity"
	"demo-dashboard/internal/log"
	"demo-dashboard/internal/utils"
	"demo-dashboard/internal/utils/gvalidator"
	"fmt"
	"strings"
	"time"

	"github.com/gofiber/fiber/v2"
)

/*
var (
	client = resty.New()
	url = fmt.Sprintf("%s/apisix/admin/upstreams", conf.ApisixConfig.AdminAPI)
)
*/

type Page struct {
	PageSize int `json:"page_size" validate:"required,page_size" default:"10"`
	PageNum  int `json:"page_num" validate:"required,page_num,min=1" default:"1"`
}

func Get(c *fiber.Ctx) error {
	appFiber := handler.Fiber{C: c}

	id := c.Params("id")
	if id == "" {
		return appFiber.Handler(fiber.StatusBadRequest, fiber.StatusBadRequest, "请求缺少ID", nil)
	}
	res, err := models.GetUpstreamByID(id)
	if err != nil {
		return appFiber.Handler(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "data error", nil)
	}
	return appFiber.Handler(fiber.StatusOK, fiber.StatusOK, utils.Success, res)
}

func GetList(c *fiber.Ctx) error {
	appFiber := handler.Fiber{C: c}
	page := Page{
		PageSize: c.QueryInt("page_size"),
		PageNum:  c.QueryInt("page_num"),
	}
	xvalidator := gvalidator.New()
	errs := xvalidator.ErrMsgs(page)
	if len(errs) > 0 {
		return appFiber.Handler(fiber.StatusBadRequest, fiber.StatusBadRequest, strings.Join(errs, " and "), nil)
	}

	res, err := models.GetUpstreamList(page.PageNum, page.PageSize)
	if err != nil {
		return appFiber.Handler(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "获取数据列表失败", nil)
	}
	return appFiber.Handler(fiber.StatusOK, fiber.StatusOK, utils.Success, res)
}

func Post(c *fiber.Ctx) error {
	appFiber := handler.Fiber{C: c}
	upstream := &entity.Upstream{}
	if err := c.BodyParser(upstream); err != nil {
		log.Logger.Error(err)
		return appFiber.Handler(fiber.StatusBadRequest, fiber.StatusBadRequest, fiber.ErrBadRequest.Error(), nil)
	}
	xvalidator := gvalidator.New()
	errs := xvalidator.ErrMsgs(upstream)
	if len(errs) > 0 {
		return appFiber.Handler(fiber.StatusBadRequest, fiber.StatusBadRequest, strings.Join(errs, " and "), nil)
	}

	fagent := fiber.Post(entity.UpstreamURL)
	fagent.Set("X-API-KEY", conf.ApisixConfig.Token)
	fagent.Timeout(5 * time.Second)
	fagent.Body(c.Body())
	code, body, err := fagent.String()
	if err != nil {
		return c.Status(code).JSON(fiber.Map{
			"errs": err,
		})
	}
	return appFiber.Handler(code, code, "success", body)
}

func Put(c *fiber.Ctx) error {
	appFiber := handler.Fiber{C: c}
	if c.Params("id") == "" {
		return appFiber.Handler(fiber.StatusBadRequest, fiber.StatusBadRequest, fiber.ErrBadRequest.Error(), nil)
	}
	upstream := &entity.Upstream{}
	if err := c.BodyParser(upstream); err != nil {
		log.Logger.Error(err)
		return appFiber.Handler(fiber.StatusBadRequest, fiber.StatusBadRequest, fiber.ErrBadRequest.Error(), nil)
	}
	xvalidator := gvalidator.New()
	errs := xvalidator.ErrMsgs(upstream)
	if len(errs) > 0 {
		return appFiber.Handler(fiber.StatusBadRequest, fiber.StatusBadRequest, strings.Join(errs, " and "), nil)
	}

	fagent := fiber.Put(fmt.Sprintf("%s/%s", entity.UpstreamURL, c.Params("id")))
	fagent.Set("X-API-KEY", conf.ApisixConfig.Token)
	fagent.Timeout(5 * time.Second)
	fagent.Body(c.Body())
	code, body, err := fagent.String()
	if err != nil {
		return c.Status(code).JSON(fiber.Map{
			"errs": err,
		})
	}
	return appFiber.Handler(code, code, "success", body)
}

func Delete(c *fiber.Ctx) error {
	appFiber := handler.Fiber{C: c}
	if c.Params("id") == "" {
		return appFiber.Handler(fiber.StatusBadRequest, fiber.StatusBadRequest, fiber.ErrBadRequest.Error(), nil)
	}
	fagent := fiber.Delete(fmt.Sprintf("%s/%s", entity.UpstreamURL, c.Params("id")))
	fagent.Set("X-API-KEY", conf.ApisixConfig.Token)
	fagent.Timeout(5 * time.Second)
	code, body, err := fagent.String()
	if err != nil {
		return c.Status(code).JSON(fiber.Map{
			"errs": err,
		})
	}
	return appFiber.Handler(code, code, "success", body)
}

func Patch(c *fiber.Ctx) error {

	return nil
}

func Search(c *fiber.Ctx) error {
	appFiber := handler.Fiber{C: c}

	if c.Params("name") == "" {
		return appFiber.Handler(fiber.StatusBadRequest, fiber.StatusBadRequest, "缺少请求参数name", nil)
	}
	res, err := models.GetUpstreamByName(c.Params("name"))
	if err != nil {
		return appFiber.Handler(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "获取数据列表失败", nil)
	}

	return appFiber.Handler(fiber.StatusOK, fiber.StatusOK, "success", res)
}

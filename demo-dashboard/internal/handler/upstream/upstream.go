package upstream

import (
	"demo-dashboard/internal/conf"
	"demo-dashboard/internal/handler"
	"demo-dashboard/internal/log"
	"fmt"

	"github.com/go-resty/resty/v2"
	"github.com/gofiber/fiber/v2"
)

var (
	client = resty.New()
	url    = fmt.Sprintf("%s/apisix/admin/upstreams", conf.ApisixConfig.AdminAPI)
)

func Get(c *fiber.Ctx) error {
	appFiber := handler.Fiber{C: c}

	if err := c.Query("id"); err == "" {
		return appFiber.Handler(fiber.StatusBadRequest, fiber.StatusBadRequest, "缺少必须的upstream id", nil)
	}
	getUrl := fmt.Sprintf("%s/%s", url, c.Query("id"))
	resp, err := client.R().Get(getUrl)
	if err != nil {
		log.Logger.Error(err)
		return appFiber.Handler(fiber.StatusInternalServerError, fiber.StatusInternalServerError, "获取数据失败", nil)
	}

	log.Logger.Info(resp)
	return nil
}

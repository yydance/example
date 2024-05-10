package resty

import (
	"demo-dashboard/internal/conf"

	"github.com/go-resty/resty/v2"
)

var (
	RR = resty.New().R().SetHeader("X-API-KEY", conf.ApisixConfig.Token)
)

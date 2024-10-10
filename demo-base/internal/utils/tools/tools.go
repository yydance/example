package tools

import (
	"demo-base/internal/utils/logger"
	"strconv"
)

func StrToUint(str string) uint {

	res, err := strconv.ParseUint(str, 10, 64)
	if err != nil {
		logger.Panicf("StrToUint error: %v", err)
		//return 0
	}
	return uint(res)
}

package tools

import (
	"demo-base/internal/utils/logger"
	"reflect"
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

func StrToInt(str string) int {
	res, err := strconv.Atoi(str)
	if err != nil {
		logger.Panicf("StrToInt error: %v", err)
		//return 0
	}
	return res
}

func StringSliceToIntergerSlice(s []string) []interface{} {
	res := make([]interface{}, 0, len(s))
	for i, str := range s {
		res[i] = str
	}
	return res
}

func DiffSlices(a, b [][]string) [][]string {
	result := make([][]string, 0)
	if len(a) < len(b) {
		return DiffSlices(b, a)
	}
	for _, itemA := range a {
		found := false
		for _, itemB := range b {
			if reflect.DeepEqual(itemA, itemB) {
				found = true
				break
			}
		}
		if !found {
			result = append(result, itemA)
		}
	}
	return result
}

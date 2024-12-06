package pkg

import (
	"fmt"
	"strconv"
)

func ApisixK8sSvcName(namespace, name string, portName any) string {
	switch v := portName.(type) {
	case string:
		portName = v
	case int32:
		portName = strconv.Itoa(int(v))
	}
	return fmt.Sprintf("%s/%s:%s", namespace, name, portName)
}

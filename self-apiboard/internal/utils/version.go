package utils

import (
	"fmt"
	"os"
	"self-apiboard/internal/conf"
)

var version = conf.Version

func PrintVersion() {
	fmt.Fprintf(os.Stdout, "%-8s: %s\n", "Version", version)
}

package utils

import (
	"demo-dashboard/internal/conf"
	"fmt"
	"os"
)

var version = conf.Version

func PrintVersion() {
	fmt.Fprintf(os.Stdout, "%-8s: %s\n", "Version", version)
}

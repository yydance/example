package cmd

import (
	"demo-dashboard/internal/utils"

	"github.com/spf13/cobra"
)

func newVersionCommand() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "show manager-api version",
		Run: func(cmd *cobra.Command, args []string) {
			utils.PrintVersion()
		},
	}
}

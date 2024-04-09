package cmd

import (
	"fmt"

	"self-apiboard/internal/conf"
	"self-apiboard/internal/log"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "self-apiboard",
	Short: "APISIX dashboard based on offical API",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := manageAPI()
		return err
	},
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&conf.ConfigFile, "config", "c", "", "config file")
	rootCmd.PersistentFlags().StringVarP(&conf.WorkDir, "work-dir", "p", ".", "currect work directory")

	rootCmd.AddCommand(
		newVersionCommand(),
	)
}

func Execute() {
	fmt.Println("")
}

func manageAPI() error {
	conf.InitConf()
	log.InitLogger()

	return nil
}

package cmd

import (
	"fmt"
	"os"

	"demo-dashboard/internal/conf"
	"demo-dashboard/internal/log"
	"demo-dashboard/internal/routers"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apisix dashboard",
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
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}

func manageAPI() error {
	conf.InitConf()
	log.InitLogger()

	app := routers.InitRouter()
	app.Listen(fmt.Sprintf(":%d", conf.ServerPort))

	return nil
}

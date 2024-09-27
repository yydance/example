package cmd

import (
	"demo-base/internal/conf"
	"demo-base/internal/routers"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "demo-base",
	Short: "A demo based on gofiber",
	RunE: func(cmd *cobra.Command, args []string) error {
		err := mainCmd()
		return err
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&conf.ConfigFile, "config", "c", "config file (default is ./conf/config.yaml)")
	//rootCmd.PersistentFlags().StringVar(&conf.WorkDir, "workdir", "w", "workdir path (default is ./)")
	rootCmd.AddCommand(&cobra.Command{
		Use:   "version",
		Short: "show app version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Fprintf(os.Stdout, "version: %s\n", conf.Version)
		},
	})
}

func mainCmd() error {
	conf.InitConfig()

	errSig := make(chan error, 2)
	app := routers.InitRouter()
	err := app.Listen(fmt.Sprintf(":%d", conf.ServerConfig.Listen.Port))
	if err != nil {
		errSig <- err
	}
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-quit:
		app.ShutdownWithTimeout(30 * time.Second)
		fmt.Println("app shutdown")
	case err := <-errSig:
		fmt.Printf("app start error: %s", err)
		return err
	}
	return nil
}

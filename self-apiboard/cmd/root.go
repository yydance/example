package cmd

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"self-apiboard/internal/conf"
	"self-apiboard/internal/core/server"
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
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
	}
}

func manageAPI() error {
	conf.InitConf()
	log.InitLogger()

	s, err := server.NewServer(&server.Options{})
	if err != nil {
		return err
	}

	errSig := make(chan error, 5)
	s.Start(errSig)

	// check mysql

	// Signal received to the process externally
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		log.Infof("the server receieve %s and start shutting down", sig.String())
		//停止mysql check
		s.Stop()
		log.Infof("See you next time!")
	case err := <-errSig:
		log.Errorf("The server start failed: %s", err.Error())
		return err
	}
	return nil
}

/*
func mysqlConnectionChecker() context.CancelFunc {
	ctx, cancel := context.WithCancel(context.TODO())

	return cancel
}
*/

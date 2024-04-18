package cmd

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"demo-dashboard/internal/conf"
	"demo-dashboard/internal/log"
	"demo-dashboard/internal/routers"
	"demo-dashboard/internal/storage"

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

	errSig := make(chan error, 5)
	app := routers.InitRouter()
	err := app.Listen(fmt.Sprintf(":%d", conf.ServerPort))
	if err != nil {
		errSig <- err
	}

	stopEtcdConnectionChecker := etcdConnectionChecker()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	select {
	case sig := <-quit:
		log.Logger.Infof("The Manager API server receive %s and start shutting down", sig.String())
		stopEtcdConnectionChecker()
		app.ShutdownWithTimeout(30 * time.Second)
		log.Logger.Info("Good luck, see you next time!")
	case err := <-errSig:
		log.Logger.Errorf("The Manager API start failed: %s", err.Error())
		return err
	}
	return nil
}

func etcdConnectionChecker() context.CancelFunc {
	ctx, cancel := context.WithCancel(context.TODO())
	unavailableTimes := 0

	go func() {
		etcdClient := storage.EtcdStorageV3.Conn()
		for {
			select {
			case <-time.Tick(10 * time.Second):
				sCtx, sCancel := context.WithTimeout(ctx, 5*time.Second)
				err := etcdClient.Sync(sCtx)
				sCancel()
				if err == nil {
					unavailableTimes = 0
				}
				if err != nil {
					unavailableTimes++
					log.Logger.Errorf("etcd connection loss detected, times: %d", unavailableTimes)
					continue
				}
			case <-ctx.Done():
				return
			}
		}
	}()

	return cancel
}

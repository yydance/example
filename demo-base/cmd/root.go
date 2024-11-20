package cmd

import (
	"context"
	"demo-base/internal/conf"
	"demo-base/internal/models"
	"demo-base/internal/routers"
	"demo-base/internal/service"
	"demo-base/internal/utils/logger"
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
	rootCmd.PersistentFlags().StringVarP(&conf.ConfigFile, "config", "c", "./conf/config.yaml", "config file (default is ./conf/config.yaml)")
	rootCmd.PersistentFlags().StringVarP(&conf.WorkDir, "workdir", "w", ".", "workdir path (default is ./)")
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
	models.InitStorage()

	printConfig()
	models.EtcdStorage.Init()
	service.InitAdmin()

	errSig := make(chan error, 2)
	app := routers.InitRouter()
	go func() {
		err := app.Listen(fmt.Sprintf(":%s", conf.ServerConfig.Listen.Port))
		if err != nil {
			errSig <- err
		}
	}()

	stopEtcdConnectionChecker := etcdCheck()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT)

	select {
	case <-quit:
		logger.Info("app shutdown...")
		stopEtcdConnectionChecker()
		app.ShutdownWithTimeout(30 * time.Second)
	case err := <-errSig:
		logger.Errorf("app start error: %s", err)
		return err
	}
	return nil
}

func etcdCheck() context.CancelFunc {

	ctx, cancle := context.WithCancel(context.TODO())
	unavailable := 0

	go func() {
		etcdClient := models.EtcdStorage.Client
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.Tick(10 * time.Second):
				sCtx, sCancel := context.WithTimeout(ctx, 5*time.Second)
				err := etcdClient.Sync(sCtx)
				sCancel()
				if err != nil {
					unavailable++
					logger.Errorf("etcd connection loss detected, unavailable count: %d", unavailable)
					continue
				}
				if unavailable >= 1 {
					logger.Warnf("etcd connection recovered, unavailable count: %d", unavailable)
					unavailable = 0
					// TODO: 重载etcd中的key/value到内存中
				} else {
					logger.Debug("etcd connection is healthy")
				}
			}
		}
	}()

	return cancle
}

func printConfig() {
	// 输出配置信息，监听地址端口、日志级别、数据库配置（MySQL、etcd地址端口)等
	msg := fmt.Sprintf("listen: %s, port: %s, log level: %s, mysql: %s:%d, etcd: %s",
		conf.ServerConfig.Listen.Host,
		conf.ServerConfig.Listen.Port,
		conf.LogLevel,
		conf.MysqlConfig.Host,
		conf.MysqlConfig.Port,
		conf.EtcdConfig.Endpoints,
	)
	logger.Infof("config, %s", msg)
}

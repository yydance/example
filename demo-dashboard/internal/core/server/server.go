package server

import (
	"context"
	"demo-dashboard/internal/conf"
	"demo-dashboard/internal/log"
	"demo-dashboard/internal/utils"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
)

type server struct {
	server  *fiber.App
	options *conf.ServerConfig
}

func NewServer(options *conf.ServerConfig) (*server, error) {
	return &server{options: options}, nil
}

func (s *server) Start(errSig chan error) {
	err := s.init()
	if err != nil {
		errSig <- err
		return
	}
	s.printInfo()

	log.Infof("The apisix dashboard is listening on %s", conf.ServerHost)
	go func() {
		err := s.server.Listen(conf.ServerHost + ":" + strconv.Itoa(conf.ServerPort))
		if err != nil {
			log.Errorf("failed to listen: %s", err)
			errSig <- err
		}
	}()
}

func (s *server) Stop() {
	ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
	defer cancel()
	s.server.ShutdownWithContext(ctx)
}

func (s *server) init() error {
	log.Info("Initialize apisix manager dashboard server")
	s.setupAPI()

	return nil
}

func (s *server) printInfo() {
	fmt.Fprint(os.Stdout, "The manager-api is running successfully!\n\n")
	utils.PrintVersion()
	fmt.Fprintf(os.Stdout, "%-8s: %s\n", "Config File", viper.ConfigFileUsed())
	fmt.Fprintf(os.Stdout, "%-8s: %s:%d\n", "Listen", conf.ServerHost, conf.ServerPort)
	fmt.Fprintf(os.Stdout, "%-8s: %s\n", "Loglevel", conf.ErrorLogLevel)
	fmt.Fprintf(os.Stdout, "%-8s: %s\n", "ErrorLogFile", conf.ErrorLogPath)
	fmt.Fprintf(os.Stdout, "%-8s: %s\n\n", "AccessLogFile", conf.AccessLogPath)
}

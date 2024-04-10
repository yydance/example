package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"self-apiboard/internal/conf"
	"self-apiboard/internal/log"
	"self-apiboard/internal/utils"
	"time"

	"github.com/spf13/viper"
)

type server struct {
	server  *http.Server
	options *Options
}

type Options struct{}

func NewServer(options *Options) (*server, error) {
	return &server{options: options}, nil
}

func (s *server) Start(errSig chan error) {
	err := s.init()
	if err != nil {
		errSig <- err
		return
	}
	s.printInfo()

	log.Infof("The apisix dashboard is listening on %s", s.server.Addr)
	go func() {
		err := s.server.ListenAndServe()
		if err != nil && err != http.ErrServerClosed {
			log.Errorf("listen and serv fail: %s", err)
			errSig <- err
		}
	}()
}

func (s *server) Stop() {
	s.shutdownServer(s.server)
}

func (s *server) init() error {
	log.Info("Initialize apisix manager dashboard server")
	s.setupAPI()

	return nil
}

func (s *server) shutdownServer(server *http.Server) {
	if server != nil {
		ctx, cancel := context.WithTimeout(context.TODO(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			log.Errorf("shutting down server error: %s", err)
		}
	}
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

package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/viper"
)

const (
	EnvDEV  = "dev"
	EnvPROD = "prod"
)

var (
	ENV           string
	MysqlConfig   *Mysql
	ApisixConfig  *Apisix
	ServerOption  *ServerConfig
	Version                     = "3.8.0"
	ConfigFile                  = ""
	ServerHost                  = "0.0.0.0"
	ServerPort                  = 80
	WorkDir                     = "."
	ErrorLogLevel               = "warn"
	ErrorLogPath                = "logs/error.log"
	AccessLogPath               = "logs/access.log"
	Timeout       time.Duration = 60 * time.Second
)

type Listen struct {
	Host string
	Port int
}

type ErrorLog struct {
	Level    string
	FilePath string `json:"file_path"`
}

type AccessLog struct {
	FilePath string `json:"file_path"`
}

type Log struct {
	ErrorLog  ErrorLog
	AccessLog AccessLog
}

type ServerConfig struct {
	AppName         string        `json:"app_name,omitempty"`
	BodyLimit       int           `json:"body_limit,omitempty"`
	Concurrency     int           `json:"concurrency,omitempty"`
	IdleTimeout     time.Duration `json:"idle_timeout,omitempty"`
	Network         string        `json:"network,omitempty"`
	Prefork         bool          `json:"prefork,omitempty"`
	ReadBufferSize  int           `json:"read_buffer_size,omitempty"`
	ReadTimeout     time.Duration `json:"read_timeout,omitempty"`
	WriteBufferSize int           `json:"write_buffer_size,omitempty"`
	WriteTimeout    time.Duration `json:"write_timeout,omitempty"`
}

type Main struct {
	Listen       Listen
	Log          Log
	ServerConfig ServerConfig
}

type Apisix struct {
	AdminAPI   string `json:"admin_api"`
	ControlAPI string `json:"control_api"`
	Token      string `json:"token"`
}

type Mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
}

type Config struct {
	Main   Main
	Apisix Apisix
	Mysql  Mysql
}

func InitConf() {
	setupConfig()
	setupEnv()
}

func setupConfig() {
	if ConfigFile == "" {
		ConfigFile = "conf.yaml"
		if profile := os.Getenv("RUN_PROFILE"); profile != "" {
			ConfigFile = "conf" + "-" + profile + ".yaml"
		}
		viper.SetConfigName(ConfigFile)
		viper.SetConfigType("yaml")
		viper.AddConfigPath(WorkDir + "/conf")
	} else {
		viper.SetConfigFile(ConfigFile)
	}

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("fail to read configuration, err: %s", err.Error()))
	}

	viper.WatchConfig()

	config := Config{}
	err := viper.Unmarshal(&config)
	if err != nil {
		panic(fmt.Sprintf("fail to unmarshall configuration: %s, err: %s", ConfigFile, err.Error()))
	}

	if config.Main.Listen.Host != "" {
		ServerHost = config.Main.Listen.Host
	}
	if config.Main.Listen.Port != 0 {
		ServerPort = config.Main.Listen.Port
	}

	if config.Main.Log.ErrorLog.Level != "" {
		ErrorLogLevel = config.Main.Log.ErrorLog.Level
	}
	if config.Main.Log.AccessLog.FilePath != "" {
		AccessLogPath = config.Main.Log.AccessLog.FilePath
	}
	if config.Main.Log.ErrorLog.FilePath != "" {
		ErrorLogPath = config.Main.Log.ErrorLog.FilePath
	}

	if config.Apisix.AdminAPI == "" {
		panic("Not found apisix admin api")
	}
	if config.Apisix.Token == "" {
		panic("Not found apisix admin token")
	}

	if !filepath.IsAbs(ErrorLogPath) {
		// 这里没有做windows路径判断
		ErrorLogPath, err = filepath.Abs(filepath.Join(WorkDir, ErrorLogPath))
		if err != nil {
			panic(err)
		}
	}
	if !filepath.IsAbs(AccessLogPath) {
		// 没有做windows路径判断
		AccessLogPath, err = filepath.Abs(filepath.Join(WorkDir, AccessLogPath))
		if err != nil {
			panic(err)
		}
	}

	ApisixConfig = &config.Apisix
	MysqlConfig = &config.Mysql
	ServerOption = &config.Main.ServerConfig
}

func setupEnv() {
	// viper.AutomaticEnv()
	// ENV = viper.Get("ENV")
	ENV = EnvPROD
	if env := os.Getenv("ENV"); env != "" {
		ENV = env
	}
}

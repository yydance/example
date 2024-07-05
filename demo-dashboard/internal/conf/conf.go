package conf

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
)

const (
	EnvDEV  = "dev"
	EnvPROD = "prod"
)

var (
	ENV           string
	MysqlConfig   Mysql
	ApisixConfig  Apisix
	ServerOption  ServerConfig
	ETCDConfig    *Etcd
	Version                     = "3.8.0"
	ConfigFile                  = ""
	ServerHost                  = "0.0.0.0"
	ServerPort                  = 80
	WorkDir                     = "."
	ErrorLogLevel               = "warn"
	ErrorLogPath                = "logs/error.log"
	AccessLogPath               = "logs/access.log"
	Timeout       time.Duration = 60 * time.Second
	Jwt           JWT
)

type Listen struct {
	Host string
	Port int
}

type ErrorLog struct {
	Level    string
	FilePath string `mapstructure:"file_path"`
}

type AccessLog struct {
	FilePath string `mapstructure:"file_path"`
}

type Log struct {
	ErrorLog  ErrorLog
	AccessLog AccessLog
}

type ServerConfig struct {
	AppName         string        `mapstructure:"app_name"`
	BodyLimit       int           `mapstructure:"body_limit"`
	Concurrency     int           `mapstructure:"concurrency"`
	IdleTimeout     time.Duration `mapstructure:"idle_timeout"`
	Network         string        `mapstructure:"network"`
	Prefork         bool          `mapstructure:"prefork"`
	ReadBufferSize  int           `mapstructure:"read_buffer_size"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteBufferSize int           `mapstructure:"write_buffer_size"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
}

type JWT struct {
	Expired int    `mapstructure:"expired"`
	Secret  string `mapstructure:"secret"`
}

type Main struct {
	Listen       Listen
	Log          Log
	ServerConfig ServerConfig `mapstructure:"server_config"`
	Jwt          JWT
}

type Apisix struct {
	AdminAPI   string `mapstructure:"admin_api"`
	ControlAPI string `mapstructure:"control_api"`
	Token      string `mapstructure:"token"`
}

type Mysql struct {
	Host     string
	Port     int
	Username string
	Password string
	DB       string
}

type MTLS struct {
	CaFile   string `mapstructure:"ca_file"`
	KeyFile  string `mapstructure:"key_file"`
	CertFile string `mapstructure:"cert_file"`
}

type Etcd struct {
	Endpoints []string
	Username  string
	Password  string
	Prefix    string
	MTLS      *MTLS
}

type Config struct {
	Main   Main
	Apisix Apisix
	Mysql  Mysql
	Etcd   Etcd
}

func InitConf() {
	setupConfig()
	setupEnv()
}

func setupConfig() {
	if ConfigFile == "" {
		ConfigFile = "conf.yaml"
		if profile := os.Getenv("RUN_PROFILE"); profile != "" {
			ConfigFile = "conf" + "." + profile + ".yaml"
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

	if len(config.Etcd.Endpoints) > 0 {
		initEtcdConfig(config.Etcd)
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
	config.Apisix.AdminAPI = strings.TrimSuffix(config.Apisix.AdminAPI, "/")

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

	ApisixConfig = config.Apisix
	MysqlConfig = config.Mysql
	ServerOption = config.Main.ServerConfig
	Jwt = config.Main.Jwt
}

func initEtcdConfig(conf Etcd) {
	var endpoints = []string{"127.0.0.1:2379"}
	if len(conf.Endpoints) > 0 {
		endpoints = conf.Endpoints
	}

	prefix := "/apisix"
	if len(conf.Prefix) > 0 {
		prefix = conf.Prefix
	}

	ETCDConfig = &Etcd{
		Endpoints: endpoints,
		Username:  conf.Username,
		Password:  conf.Password,
		MTLS:      conf.MTLS,
		Prefix:    prefix,
	}
}

func setupEnv() {
	// viper.AutomaticEnv()
	// ENV = viper.Get("ENV")
	ENV = EnvPROD
	if env := os.Getenv("ENV"); env != "" {
		ENV = env
	}
}

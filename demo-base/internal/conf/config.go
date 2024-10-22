package conf

import (
	"fmt"
	"time"

	"github.com/spf13/viper"
)

var (
	RunMode    = ""
	ConfigFile = "./conf/config.yaml"
	RBACModel  = "./conf/rbac_model.conf"
	RBACPolicy = "./conf/policy.json"
	//WorkDir                  = "."
	Timeout      time.Duration = 10 * time.Second
	Version                    = "0.0.1"
	MysqlConfig  Mysql
	FiberConfig  FiberConf
	ServerConfig Server
	CorsConfig   Cors
	LogConfig    Log
)

type Config struct {
	Server   Server   `mapstructure:"server"`
	Database Database `mapstructure:"database"`
}

type Server struct {
	Listen      Listen    `mapstructure:"listen"`
	Log         Log       `mapstructure:"log"`
	JWT         JWT       `mapstructure:"jwt"`
	Cors        Cors      `mapstructure:"cors"`
	FiberConfig FiberConf `mapstructure:"fiber_config"`
	RunMode     string    `mapstructure:"run_mode"`
}

type Listen struct {
	Host string `mapstructure:"host"`
	Port string `mapstructure:"port"`
}

type Log struct {
	Level      string `mapstructure:"level"`
	ErrorPath  string `mapstructure:"error_path"`
	AccessPath string `mapstructure:"access_path"`
}

type JWT struct {
	Secret  string `mapstructure:"secret"`
	Expired int    `mapstructure:"expired"`
}

type Cors struct {
	Enabled          bool   `mapstructure:"enabled"`
	AllowOrigins     string `mapstructure:"allow_origins"`
	AllowMethods     string `mapstructure:"allow_methods"`
	AllowHeaders     string `mapstructure:"allow_headers"`
	ExposeHeaders    string `mapstructure:"expose_headers"`
	AllowCredentials bool   `mapstructure:"allow_credentials"`
	MaxAge           int    `mapstructure:"max_age"`
}

type FiberConf struct {
	AppName         string        `mapstructure:"app_name"`
	BodyLimit       int           `mapstructure:"body_limit"`
	Concurrent      int           `mapstructure:"concurrent"`
	Network         string        `mapstructure:"network"`
	Prefork         bool          `mapstructure:"prefork"`
	ReadBufferSize  int           `mapstructure:"read_buffer_size"`
	WriteBufferSize int           `mapstructure:"write_buffer_size"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	IdleTimeout     time.Duration `mapstructure:"idle_timeout"`
}

type Database struct {
	Mysql Mysql `mapstructure:"mysql"`
}

type Mysql struct {
	Host         string        `mapstructure:"host"`
	Port         int           `mapstructure:"port"`
	User         string        `mapstructure:"user"`
	Password     string        `mapstructure:"password"`
	Db           string        `mapstructure:"db"`
	MaxIdleConns int           `mapstructure:"max_idle_connections"`
	MaxOpenConns int           `mapstructure:"max_open_connections"`
	MaxLifeTime  time.Duration `mapstructure:"max_life_time"`
	MaxIdleTime  time.Duration `mapstructure:"max_idle_time"`
}

func InitConfig() {
	setupConfig()
}

func setupConfig() {
	if ConfigFile == "" {
		ConfigFile = "./conf/config.yaml"
	}
	viper.SetConfigFile(ConfigFile)

	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Sprintf("Failed to read the configuration file: %s", err.Error()))
	}
	viper.WatchConfig()

	config := &Config{}
	if err := viper.Unmarshal(config); err != nil {
		panic(fmt.Sprintf("Failed to unmarshal the configuration file: %s, err: %s", ConfigFile, err.Error()))
	}
	MysqlConfig = config.Database.Mysql
	RunMode = config.Server.RunMode
	FiberConfig = config.Server.FiberConfig
	CorsConfig = config.Server.Cors
	LogConfig = config.Server.Log
	ServerConfig = config.Server
}

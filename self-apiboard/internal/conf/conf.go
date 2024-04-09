package conf

import (
	"fmt"
	"os"
	"path/filepath"

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
	Version       = "3.8.0"
	ConfigFile    = ""
	ServerHost    = "0.0.0.0"
	ServerPort    = 80
	WorkDir       = "."
	ErrorLogLevel = "warn"
	ErrorLogPath  = "logs/error.log"
	AccessLogPath = "logs/access.log"
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

type Main struct {
	Listen Listen
	Log    Log
}

type Apisix struct {
	AdminAPI   string `mapstructure:"admin_api"`
	ControlAPI string `mapstructure:"control_api"`
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

}

func setupEnv() {
	// viper.AutomaticEnv()
	// ENV = viper.Get("ENV")
	ENV = EnvPROD
	if env := os.Getenv("ENV"); env != "" {
		ENV = env
	}
}

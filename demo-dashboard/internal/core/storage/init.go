package storage

import (
	"demo-dashboard/internal/conf"
	"time"

	"github.com/gofiber/storage/etcd/v2"
	"github.com/gofiber/storage/mysql/v2"
)

var (
	EtcdStorageV3 *etcd.Storage
	MysqlStorage  *mysql.Storage
)

func init() {
	EtcdStorageV3 = etcd.New(etcd.Config{
		Endpoints:   conf.ETCDConfig.Endpoints,
		DialTimeout: 10 * time.Second,
		Username:    conf.ETCDConfig.Username,
		Password:    conf.ETCDConfig.Password,
	})

	MysqlStorage = mysql.New(mysql.Config{
		Host:     conf.MysqlConfig.Host,
		Port:     conf.MysqlConfig.Port,
		Database: conf.MysqlConfig.DB,
		Username: conf.MysqlConfig.Username,
		Password: conf.MysqlConfig.Password,
	})
}

func initMysqlConfig(table string) mysql.Config {
	cfg := mysql.Config{
		Host:     conf.MysqlConfig.Host,
		Port:     conf.MysqlConfig.Port,
		Database: conf.MysqlConfig.DB,
		Username: conf.MysqlConfig.Username,
		Password: conf.MysqlConfig.Password,
		Table:    table,
		Reset:    false,
	}

	return cfg
}

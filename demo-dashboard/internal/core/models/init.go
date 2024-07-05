package models

import (
	"context"
	"demo-dashboard/internal/conf"
	"demo-dashboard/internal/log"

	"fmt"
	"time"

	"github.com/gofiber/storage/etcd/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	EtcdStorageV3 *etcd.Storage
	Db            *gorm.DB
)

func InitStorage() {

	EtcdStorageV3 = etcd.New(etcd.Config{
		Endpoints:   conf.ETCDConfig.Endpoints,
		DialTimeout: 10 * time.Second,
		Username:    conf.ETCDConfig.Username,
		Password:    conf.ETCDConfig.Password,
	})

	Db = initMysqlDB(conf.MysqlConfig)
	Db.AutoMigrate(
		&UpstreamTLS{},
		&UpstreamKeepalivePool{},
		&Timeout{},
		&Upstream{},
		&Service{},
		//&Route{},
		&Consumer{},
		&GlobalPlugins{},
		&ServerInfo{},
	)
}

func initMysqlDB(conf conf.Mysql) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", conf.Username, conf.Password, conf.Host, conf.Port, conf.DB)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 30,
	})
	if err != nil {
		log.Logger.Errorf("failed to connect database: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		log.Logger.Panic(err)
	}
	sqlDB.SetConnMaxIdleTime(30 * time.Second)
	sqlDB.SetConnMaxLifetime(time.Hour)
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)

	return db
}

func initCtx(timeOut time.Duration) *gorm.DB {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	tx := Db.WithContext(ctx)
	return tx
}

func CloseDB() {
	sqlDB, err := Db.DB()
	if err != nil {
		log.Logger.Panic(err)
	}
	defer sqlDB.Close()
}

/*
func JsonParser(dType string, data []byte) any {
	var (
		create_time int64
		update_time int64
		id          any
		name        string
		utype       string
		hosts []string
		uris []string
		status bool
		desc string
	)
	_, err := jsonparser.ArrayEach(data, func(value []byte, dataType jsonparser.ValueType, offset int, err error) {
		create_time, _ = jsonparser.GetInt(value, "create_time")
		update_time, _ = jsonparser.GetInt(data, "update_time")
		id, _ = jsonparser.GetString(data, "id")
		name, _ = jsonparser.GetString(data, "name")
		switch dType {
		case "upstream":
			utype, _ = jsonparser.GetString(data, "type")
		case "route":
			status, _ = jsonparser.GetBoolean(data, "status")
			uris, _ = jsonparser.Get()
		}

	})

	if err != nil {
		log.Logger.Errorf("%s", err)
	}

	return Upstream{
		BaseInfo: BaseInfo{
			ID:         id,
			CreateTime: create_time,
			UpdateTime: update_time,
		},
		Name: name,
		Type: utype,
	}
}
*/

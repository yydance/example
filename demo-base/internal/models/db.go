package models

import (
	"demo-base/internal/conf"
	"demo-base/internal/utils/logger"
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	DB = newDB(conf.MysqlConfig)
	DB.AutoMigrate(
		&User{},
		&Role{}) // 自动迁移模式
}

func newDB(config conf.Mysql) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Password, config.Host, config.Port, config.Db)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		CreateBatchSize: 100,
	})
	if err != nil {
		logger.Errorf("failed to connect database, err: %v", err)
	}
	sqlDB, err := db.DB()
	if err != nil {
		logger.Panic(err)
	}
	sqlDB.SetConnMaxIdleTime(conf.MysqlConfig.MaxIdleTime * time.Second)
	sqlDB.SetConnMaxLifetime(conf.MysqlConfig.MaxLifeTime * time.Second)
	sqlDB.SetMaxOpenConns(conf.MysqlConfig.MaxOpenConns)
	sqlDB.SetMaxIdleConns(conf.MysqlConfig.MaxIdleConns)

	return db
}

/*
func dbCtx(timeOut time.Duration) *gorm.DB {
	ctx, cancel := context.WithTimeout(context.Background(), timeOut)
	defer cancel()
	return DB.WithContext(ctx)
}
*/

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		logger.Panic(err)
	}
	sqlDB.Close()
}

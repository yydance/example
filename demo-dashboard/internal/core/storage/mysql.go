package storage

import "github.com/gofiber/storage/mysql/v2"

type MysqlStore struct {
	Option       any
	mysqlStorage *mysql.Storage
}

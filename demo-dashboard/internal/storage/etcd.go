package storage

import (
	"demo-dashboard/internal/conf"
	"time"

	"github.com/gofiber/storage/etcd/v2"
)

type EtcdStorageV3 struct {
	db *etcd.Storage
}

func New(conf *conf.Etcd) *EtcdStorageV3 {
	db := etcd.New(etcd.Config{
		Endpoints:   conf.Endpoints,
		DialTimeout: 10 * time.Second,
		Username:    conf.Username,
		Password:    conf.Password,
	})

	store := &EtcdStorageV3{
		db: db,
	}

	return store
}

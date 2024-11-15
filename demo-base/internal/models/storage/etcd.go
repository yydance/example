package storage

import (
	"context"
	"demo-base/internal/conf"
	"demo-base/internal/utils/logger"
	"errors"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// TODO: 缺少etcd watch的context管理
const (
	SkippedValueEtcdInitDir     = "init_dir"
	SkippedValueEtcdEmptyObject = "{}"
)

type EtcdStorage struct {
	Client *clientv3.Client
	Error  error
}

func NewEtcdStorage() *EtcdStorage {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   conf.EtcdConfig.Endpoints,
		DialTimeout: conf.EtcdConfig.DialTimeout,
		Username:    conf.EtcdConfig.Username,
		Password:    conf.EtcdConfig.Password,
		TLS:         nil,
	})
	if err != nil {
		logger.Errorf("etcd client init failed: %v", err)
		return &EtcdStorage{
			Error: err,
		}
	}
	return &EtcdStorage{
		Client: client,
		Error:  nil,
	}
}

func (e *EtcdStorage) Init(ctx context.Context) error {
	err := e.Set(ctx, conf.EtcdConfig.Prefix, []byte(SkippedValueEtcdInitDir), 0)
	if err != nil {
		return errors.New("Init etcd prefix failed")
	}

	err = e.Set(ctx, conf.RolesPrefix, []byte(SkippedValueEtcdInitDir), 0)
	if err != nil {
		return errors.New("Init etcd roles prefix failed")
	}
	return nil
}

func (e *EtcdStorage) Get(ctx context.Context, key string) ([]byte, error) {
	if key == "" {
		logger.Error("key is empty")
		return nil, nil
	}
	resp, err := e.Client.Get(ctx, key)
	if err != nil {
		logger.Errorf("etcd get failed: %v", err)
		return nil, fmt.Errorf("etcd get failed: %s", err)
	}
	if len(resp.Kvs) <= 0 {
		return nil, fmt.Errorf("Key(%s) not found", key)
	}
	return resp.Kvs[0].Value, nil
}

func (e *EtcdStorage) Set(ctx context.Context, key string, value []byte, exp time.Duration) error {
	if key == "" || value == nil {
		logger.Error("key or value is empty")
		return nil
	}
	lease, err := e.Client.Grant(ctx, int64(exp.Seconds()))
	if err != nil {
		return err
	}
	_, err = e.Client.Put(ctx, key, string(value), clientv3.WithLease(lease.ID))
	if err != nil {
		logger.Errorf("etcd put failed: %v", err)
		return fmt.Errorf("etcd put failed: %s", err)
	}
	return nil
}

func (e *EtcdStorage) Delete(ctx context.Context, key string) error {
	if key == "" {
		logger.Error("key is empty")
		return nil
	}
	resp, err := e.Client.Delete(ctx, key)
	if err != nil {
		logger.Errorf("etcd delete failed: %v", err)
		return fmt.Errorf("etcd delete failed: %s", err)
	}
	if resp.Deleted == 0 {
		return fmt.Errorf("Key(%s) not found", key)
	}
	return nil
}

func (e *EtcdStorage) List(ctx context.Context, key string) ([]KeyPair, error) {
	resp, err := e.Client.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		logger.Errorf("etcd get failed: %v", err)
		return nil, fmt.Errorf("etcd get failed: %s", err)
	}
	if len(resp.Kvs) == 0 {
		return nil, fmt.Errorf("Key(%s) not found", key)
	}
	var kvs []KeyPair
	for i := range resp.Kvs {
		key := string(resp.Kvs[i].Key)
		value := string(resp.Kvs[i].Value)
		if value == SkippedValueEtcdInitDir || value == SkippedValueEtcdEmptyObject {
			continue
		}
		data := KeyPair{
			Key:   key,
			Value: value,
		}
		kvs = append(kvs, data)
	}
	return kvs, nil
}

func (e *EtcdStorage) BatchDelete(ctx context.Context, keys []string) error {
	if len(keys) == 0 {
		logger.Error("key is empty")
		return nil
	}
	for i := range keys {
		resp, err := e.Client.Delete(ctx, keys[i])
		if err != nil {
			logger.Errorf("etcd delete failed: %v", err)
			return fmt.Errorf("etcd delete failed: %s", err)
		}
		if resp.Deleted == 0 {
			return fmt.Errorf("Key(%s) not found", keys[i])
		}
	}
	return nil
}

func (e *EtcdStorage) Watch(ctx context.Context, key string) <-chan WatchResponse {
	eventChan := e.Client.Watcher.Watch(ctx, key, clientv3.WithPrefix())
	ch := make(chan WatchResponse, 1)

	go func() {
		for event := range eventChan {
			if event.Err() != nil {
				logger.Errorf("etcd watch error: key(%s), err: %v", key, event.Err())
				close(ch)
				return
			}

			output := WatchResponse{}

			for i := range event.Events {
				key := string(event.Events[i].Kv.Key)
				value := string(event.Events[i].Kv.Value)
				if value == SkippedValueEtcdInitDir || value == SkippedValueEtcdEmptyObject {
					continue
				}
				e := Event{
					KeyPair: KeyPair{
						Key:   key,
						Value: value,
					},
				}
				switch event.Events[i].Type {
				case clientv3.EventTypePut:
					e.Type = EventTypePut
				case clientv3.EventTypeDelete:
					e.Type = EventTypeDelete
				}
				output.Events = append(output.Events, e)
			}
			ch <- output
		}
		close(ch)
	}()
	return ch
}

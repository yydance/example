package models

import (
	"context"
	"demo-base/internal/conf"
	"demo-base/internal/models/storage"
	"demo-base/internal/utils/logger"
	"fmt"
	"sync"

	"github.com/bytedance/sonic"
)

type GenericStore struct {
	Cache sync.Map
}

func NewGenericStore() *GenericStore {
	return &GenericStore{}
}

func (s *GenericStore) Get(key any) (any, bool) {
	return s.Cache.Load(key)
}
func (s *GenericStore) ValueToMap(key any) (value map[string]any, err error) {
	v, ok := s.Cache.Load(key)
	if !ok {
		logger.Errorf("key %v not found", key)
		value = nil
		return
	}
	if v == nil {
		err = fmt.Errorf("key %v is nil", key)
		value = nil
		return
	}
	if err = sonic.Unmarshal([]byte(v.(string)), &value); err != nil {
		err = fmt.Errorf("key %v is not a map", key)
		value = nil
		return
	}

	return value, nil
}

// 注意：这里key是etcd key的全路径，占用内存可能较大，可以只存储key的最后一部分（即用户名）
func (s *GenericStore) InitData() error {
	keyPairs, err := EtcdStorage.List(context.TODO(), conf.RolesPrefix)
	if err != nil {
		return err
	}
	if len(keyPairs) == 0 {
		return nil
	}

	for i := range keyPairs {
		s.Cache.Store(keyPairs[i].Key, keyPairs[i].Value)

	}
	return nil
}

func (s *GenericStore) Watch() error {
	WatchResponseChan := EtcdStorage.Watch(context.TODO(), conf.RolesPrefix)
	go func() {
		for watchResponse := range WatchResponseChan {
			for i := range watchResponse.Events {
				event := watchResponse.Events[i]
				if event.Type == storage.EventTypePut {
					s.Cache.Store(event.Key, event.Value)
				}
				if event.Type == storage.EventTypeDelete {
					s.Cache.Delete(event.Key)
				}
			}
		}
	}()
	return nil
}

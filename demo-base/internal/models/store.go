package models

import (
	"context"
	"demo-base/internal/conf"
	"demo-base/internal/models/storage"
	"sync"
)

type GenericStore struct {
	Cache sync.Map
}

func (s *GenericStore) Get(key any) (any, bool) {
	return s.Cache.Load(key)
}

func (s *GenericStore) Load() error {
	keyPairs, err := EtcdStorage.List(context.TODO(), conf.RolesPrefix)
	if err != nil {
		return err
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

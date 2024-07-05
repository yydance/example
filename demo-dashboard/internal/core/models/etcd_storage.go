package models

/*
func UpstreamWatch(ctx context.Context) error {
	watcher := clientv3.NewWatcher(EtcdStorageV3.Conn())
	watchChan := watcher.Watch(ctx, conf.ETCDConfig.Prefix+"/upstreams", clientv3.WithPrefix())


	for resp := range watchChan {
		for _, ev := range resp.Events {
			key := ev.Kv.Key
			value := ev.Kv.Value
		}
	}

	return nil
}

func UpstreamInit(ctx context.Context) error {
	defer EtcdStorageV3.Close()
	basePath := fmt.Sprintf("%s/upstreams", conf.ETCDConfig.Prefix)
	resp, err := EtcdStorageV3.Conn().Get(ctx, basePath, clientv3.WithPrefix())
	if err != nil {
		log.Logger.Errorf("初始化upstreams数据失败，%v", err)
		return err
	}
	for _, kv := range resp.Kvs {
		key := string(kv.Key)[len(basePath)+1:]
		value := string(kv.Value)
		id := key

	}
	return nil
}
*/

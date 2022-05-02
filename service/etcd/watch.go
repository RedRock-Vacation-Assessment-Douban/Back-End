package etcd

import (
	"context"
	"errors"
	clientv3 "go.etcd.io/etcd/client/v3"
	"log"
	"strings"
	"time"
)

type Watcher struct {
	cli       *clientv3.Client
	Prefix    string
	WatchChan *clientv3.WatchChan
}

type KV struct {
	Key string
	Val string
}

// GetKV 可以获取所有etcd的以“/key”开头的kv
func (w *Watcher) GetKV(key string) (kvs []KV, err error) {
	get, err := w.cli.Get(context.TODO(), "/"+key, clientv3.WithPrefix())
	if err != nil {
		log.Println(err)
		return nil, err
	}
	//s := make([]string, 24)
	kvs = []KV{}
	for _, v := range get.Kvs {
		var kv KV
		kv.Key = string(v.Key)
		kv.Val = string(v.Value)
		kvs = append(kvs, kv)
	}
	return kvs, nil
}

// NewWatcher 初始化一个Watcher
func NewWatcher(endpoint string, ProjectName string) *Watcher {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{endpoint},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		log.Println(err)
		return nil
	}
	var watch Watcher
	watch.cli = client
	watch.Prefix = ProjectName
	return &watch
}

// NewWatchChan 可以返回一个*clientv3.WatchChan来检测etcd的变化，暂时只支持一个节点/key
func (w *Watcher) NewWatchChan(key ...string) (WatchChan *clientv3.WatchChan) {
	watchChan := w.cli.Watch(context.TODO(), key[0], clientv3.WithPrefix())
	w.WatchChan = &watchChan
	return &watchChan
}

// MatchChildNode 可以返回projectName的子节点的value
func (w *Watcher) MatchChildNode(child string) (values string, err error) {
	kvs, _ := w.GetKV(w.Prefix)
	for _, kv := range kvs {
		trim := strings.TrimLeft(kv.Key, "/"+w.Prefix)
		if child == trim {
			return kv.Val, nil
		}
	}
	return "", errors.New("no Child Err")
}

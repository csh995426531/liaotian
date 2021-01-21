package repository

import (
	"github.com/coreos/etcd/clientv3"
	"strings"
	"time"
)

func NewClient() (client *clientv3.Client, err error) {
	config := clientv3.Config{
		Endpoints: strings.Split("http://192.168.66.100:12379", ","),
		DialTimeout: 10*time.Second,
	}
	client, err = clientv3.New(config)
	return
}

func GetKv() clientv3.KV {
	return clientv3.NewKV(EtcdClient)
}
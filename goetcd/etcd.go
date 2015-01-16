package goetcd

import (
	"github.com/coreos/go-etcd/etcd"
)

type EtcdConfig struct {
	Client *etcd.Client
}

func NewEtcdClient(etcdaddr string) *etcd.Client {
	return etcd.NewClient([]string{etcdaddr})
}

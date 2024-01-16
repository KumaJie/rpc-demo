package etcd

import (
	"context"
	"fmt"
	clientv3 "go.etcd.io/etcd/client/v3"
	"go.etcd.io/etcd/client/v3/naming/endpoints"
	"go.uber.org/zap"
	"rpc-douyin/src/config"
	"rpc-douyin/src/util/log"
)

const (
	nameServicePrefix  = "douyin" // 服务前缀
	defaultLeaseSecond = 30       // 默认的租赁时间
)

type NamingService struct {
	Client *clientv3.Client
}

type Endpoint struct {
	Name    string
	Addr    string
	Port    int
	LeaseID clientv3.LeaseID
}

func NewNamingService() (*NamingService, error) {
	etcdURL := fmt.Sprintf("http://%s:%d", config.Cfg.Etcd.Host, config.Cfg.Etcd.Port)
	client, err := clientv3.NewFromURL(etcdURL)
	if err != nil {
		return nil, err
	}
	return &NamingService{
		Client: client,
	}, err
}

func GetServicePath(serviceName string) string {
	return fmt.Sprintf("%s/%s", nameServicePrefix, serviceName)
}

func (n *NamingService) Register(e *Endpoint) error {
	manager, err := endpoints.NewManager(n.Client, nameServicePrefix)
	if err != nil {
		return err
	}
	lease, err := n.Client.Grant(context.Background(), defaultLeaseSecond)
	if err != nil {
		return err
	}
	e.LeaseID = lease.ID
	err = manager.AddEndpoint(context.Background(), fmt.Sprintf("%s/%s/%d", nameServicePrefix, e.Name, lease.ID), endpoints.Endpoint{
		Addr:     fmt.Sprintf("%s:%d", e.Addr, e.Port),
		Metadata: nil,
	}, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}
	ch, err := n.Client.KeepAlive(context.Background(), lease.ID)
	go func() {
		for {
			<-ch
		}
	}()
	log.Info("etcd: register naming service", zap.String("name", fmt.Sprintf("%s/%s/%d", nameServicePrefix, e.Name, lease.ID)))
	return err
}

func (n *NamingService) Delete(e Endpoint) error {
	em, err := endpoints.NewManager(n.Client, nameServicePrefix)
	if err != nil {
		return err
	}
	err = em.DeleteEndpoint(context.Background(), fmt.Sprintf("%s/%s/%d", nameServicePrefix, e.Name, e.LeaseID))
	if err != nil {
		return err
	}
	_, err = n.Client.Revoke(context.Background(), e.LeaseID)
	return err
}

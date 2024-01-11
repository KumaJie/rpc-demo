package connectWrapper

import (
	"fmt"
	"go.etcd.io/etcd/client/v3/naming/resolver"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"rpc-douyin/src/config"
	"rpc-douyin/src/storage/etcd"
)

func Connect(serviceName string) *grpc.ClientConn {
	target := fmt.Sprintf("etcd://%s:%d/%s",
		config.Cfg.Etcd.Host,
		config.Cfg.Etcd.Port,
		etcd.GetServicePath(serviceName))
	service, err := etcd.NewNamingService()
	if err != nil {
		panic(err)
	}
	etcdResolver, err := resolver.NewBuilder(service.Client)
	if err != nil {
		panic(err)
	}
	conn, err := grpc.Dial(
		target,
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithResolvers(etcdResolver),
		grpc.WithDefaultServiceConfig(`{"loadBalancingPolicy":"round_robin"}`))
	if err != nil {
		fmt.Printf("connect to grpc faild: %v", err)
	}
	return conn
}

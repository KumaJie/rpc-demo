package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"rpc-douyin/src/config"
	"rpc-douyin/src/proto/auth"
	"rpc-douyin/src/storage/etcd"
	"rpc-douyin/src/util/log"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.Server.Auth.Port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	auth.RegisterAuthServiceServer(s, &AuthServiceImpl{})

	service, err := etcd.NewNamingService()
	if err != nil {
		panic(err)
	}
	e := etcd.Endpoint{
		Name: config.Cfg.Server.Auth.Name,
		Addr: config.Cfg.Server.Auth.Host,
		Port: config.Cfg.Server.Auth.Port,
	}
	err = service.Register(&e)
	if err != nil {
		panic(err)
	}
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to server: %v", err)
		return
	}
	log.Info("AuthService: service start", zap.Int("port", config.Cfg.Server.Auth.Port))
}

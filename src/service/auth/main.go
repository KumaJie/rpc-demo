package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"rpc-douyin/src/config"
	"rpc-douyin/src/proto/auth"
	"rpc-douyin/src/storage/etcd"
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
		Addr: "127.0.0.1",
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
}

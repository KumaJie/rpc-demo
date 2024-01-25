package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"rpc-douyin/src/config"
	"rpc-douyin/src/proto/user"
	"rpc-douyin/src/storage/etcd"
	"rpc-douyin/src/util/log"
	"rpc-douyin/src/util/tracer"
)

func main() {
	cfg := config.Cfg.Server.User
	tracer.InitTracer(cfg.Name, fmt.Sprintf("%s:%d", cfg.Host, cfg.Port))
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", cfg.Port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, &UserServiceImpl{})

	service, err := etcd.NewNamingService()
	if err != nil {
		panic(err)
	}
	e := etcd.Endpoint{
		Name: config.Cfg.Server.User.Name,
		Addr: config.Cfg.Server.User.Host,
		Port: config.Cfg.Server.User.Port,
	}
	err = service.Register(&e)
	if err != nil {
		panic(err)
	}
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to server: %v", err)
		service.Delete(e)
		return
	}
	log.Info("UserService: service start", zap.Int("port", cfg.Port))
}

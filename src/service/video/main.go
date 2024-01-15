package main

import (
	"fmt"
	"google.golang.org/grpc"
	"net"
	"rpc-douyin/src/config"
	"rpc-douyin/src/proto/video"
	"rpc-douyin/src/storage/etcd"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.Server.Video.Port))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	video.RegisterVideoServiceServer(s, &VideoServiceImpl{})

	service, err := etcd.NewNamingService()
	if err != nil {
		panic(err)
	}
	e := etcd.Endpoint{
		Name: config.Cfg.Server.Video.Name,
		Addr: "127.0.0.1",
		Port: config.Cfg.Server.Video.Port,
	}
	err = service.Register(&e)
	if err != nil {
		panic(err)
	}
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}

}

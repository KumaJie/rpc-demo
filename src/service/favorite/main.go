package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"rpc-douyin/src/config"
	"rpc-douyin/src/proto/favorite"
	"rpc-douyin/src/storage/etcd"
	"rpc-douyin/src/util/log"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.Server.Favorite.Port))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	favorite.RegisterFavoriteServiceServer(s, &FavoriteServiceImpl{})
	service, err := etcd.NewNamingService()
	if err != nil {
		panic(err)
	}

	e := etcd.Endpoint{
		Name: config.Cfg.Server.Favorite.Name,
		Addr: "127.0.0.1",
		Port: config.Cfg.Server.Favorite.Port,
	}
	err = service.Register(&e)
	if err != nil {
		panic(err)
	}
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
	log.Info("FavoriteService: service start", zap.Int("port", config.Cfg.Server.Comment.Port))
}

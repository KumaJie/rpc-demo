package main

import (
	"fmt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
	"rpc-douyin/src/config"
	"rpc-douyin/src/proto/comment"
	"rpc-douyin/src/storage/etcd"
	"rpc-douyin/src/util/log"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", config.Cfg.Server.Comment.Port))
	if err != nil {
		panic(err)
	}
	s := grpc.NewServer()
	comment.RegisterCommentServiceServer(s, &CommentServiceImpl{})

	service, err := etcd.NewNamingService()
	if err != nil {
		panic(err)
	}

	e := etcd.Endpoint{
		Name: config.Cfg.Server.Comment.Name,
		Addr: config.Cfg.Server.Comment.Host,
		Port: config.Cfg.Server.Comment.Port,
	}
	err = service.Register(&e)
	if err != nil {
		panic(err)
	}
	err = s.Serve(lis)
	if err != nil {
		panic(err)
	}
	log.Info("CommentService: service start", zap.Int("port", config.Cfg.Server.Comment.Port))
}

package main

import (
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"net"
	"rpc-douyin/src/proto/user"
)

var (
	port = flag.Int("port", 8000, "The server port")
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}
	s := grpc.NewServer()
	user.RegisterUserServiceServer(s, &UserServiceImpl{})
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to server: %v", err)
		return
	}
}

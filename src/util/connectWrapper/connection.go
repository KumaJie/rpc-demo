package connectWrapper

import (
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	host = "127.0.0.1"
)

func Connect(port int) *grpc.ClientConn {
	conn, err := grpc.Dial(fmt.Sprintf("%s:%d", host, port), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		fmt.Printf("connect to grpc faild: %v", err)
	}
	return conn
}

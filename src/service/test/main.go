package main

import (
	"context"
	"fmt"
	"rpc-douyin/src/proto/user"
	"rpc-douyin/src/util/connectWrapper"
)

func main() {
	conn := connectWrapper.Connect(8000)
	defer conn.Close()
	client := user.NewUserServiceClient(conn)
	resp, err := client.GetUserExist(context.Background(), &user.UserRequest{UserId: 1})
	if err != nil {
		fmt.Printf("failed: %v", err)
		return
	}
	fmt.Println(resp.StatusCode)
}

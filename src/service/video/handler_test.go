package main

import (
	"context"
	"fmt"
	"rpc-douyin/src/config"
	"rpc-douyin/src/proto/video"
	"rpc-douyin/src/util/connectWrapper"
	"testing"
	"time"
)

func TestVideoServiceImpl_Feed(t *testing.T) {
	conn := connectWrapper.Connect(config.Cfg.Server.Video.Name)
	defer conn.Close()
	client := video.NewVideoServiceClient(conn)

	t.Run("最新", func(t *testing.T) {
		latestTime := time.Now().UnixMilli()
		ret, _ := client.Feed(context.Background(), &video.FeedRequest{LatestTime: &latestTime})
		for _, v := range ret.VideoList {
			fmt.Println(v)
		}
	})
}

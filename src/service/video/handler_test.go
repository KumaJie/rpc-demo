package main

import (
	"context"
	"github.com/stretchr/testify/assert"
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
		assert.Equal(t, len(ret.GetVideoList()), 3)
		assert.Less(t, ret.GetNextTime(), latestTime)
	})

	t.Run("全部", func(t *testing.T) {
		latestTime := time.Now().UnixMilli()
		for {
			ret, err := client.Feed(context.Background(), &video.FeedRequest{LatestTime: &latestTime})
			assert.NoError(t, err)
			if len(ret.GetVideoList()) == 0 {
				assert.Zero(t, ret.GetNextTime())
				break
			}
			latestTime = ret.GetNextTime()
		}
	})
}

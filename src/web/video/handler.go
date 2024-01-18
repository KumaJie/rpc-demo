package video

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"rpc-douyin/src/config"
	"rpc-douyin/src/model"
	"rpc-douyin/src/proto/video"
	"rpc-douyin/src/util/connectWrapper"
	"strconv"
	"time"
)

var videoClient video.VideoServiceClient

func init() {
	conn := connectWrapper.Connect(config.Cfg.Server.Video.Name)
	videoClient = video.NewVideoServiceClient(conn)
}

func VideoPublishHandler(c *gin.Context) {
	userID, _ := c.Get("user_id")
	form, _ := c.MultipartForm()
	title := form.Value["title"][0]

	file := form.File["data"][0]
	fd, _ := file.Open()

	defer func(fd multipart.File) {
		fd.Close()
	}(fd)

	rawData := make([]byte, file.Size)
	fd.Read(rawData)

	_, err := videoClient.VideoPublish(context.Background(), &video.PublishRequest{
		UserId: userID.(int64),
		Data:   rawData,
		Title:  title,
	})
	if err != nil {
		fmt.Println(err)
	}
	c.JSON(http.StatusOK, &model.Response{
		StatusCode: 0,
		StatusMsg:  "上传成功",
	})
}

func PublishListHandler(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, _ := strconv.ParseInt(userIDStr, 10, 64)

	ret, _ := videoClient.GetPublishList(context.Background(), &video.PublishListRequest{UserId: userID})
	videoList := ret.GetVideoList()
	c.JSON(http.StatusOK, model.VideoResp{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		VideoList: videoList,
	})
}

func FeedHandler(c *gin.Context) {
	// 首次请求会发起两次，第一次是latest_time = now()，第二次是latest_time = next_time(第一次请求返回的结果，即最早的发布时间)
	latestTimeStr := c.Query("latest_time")
	latestTime, _ := strconv.ParseInt(latestTimeStr, 10, 64)
	fmt.Println(time.UnixMilli(latestTime))
	feedResp, _ := videoClient.Feed(context.Background(), &video.FeedRequest{LatestTime: &latestTime})
	c.JSON(http.StatusOK, model.FeedResp{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		VideoList: feedResp.GetVideoList(),
		NextTime:  feedResp.NextTime,
	})
}

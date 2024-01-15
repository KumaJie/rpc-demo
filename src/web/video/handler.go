package video

import (
	"context"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"rpc-douyin/src/config"
	"rpc-douyin/src/model"
	"rpc-douyin/src/proto/video"
	"rpc-douyin/src/util/connectWrapper"
	"strconv"
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

	}
	c.JSON(http.StatusOK, &model.Response{
		StatusCode: 0,
		StatusMsg:  "上传成功",
	})
}

func PublishListHandler(c *gin.Context) {
	userIDStr := c.Query("user_id")
	userID, _ := strconv.ParseInt(userIDStr, 10, 64)

	ret, _ := videoClient.GetPublistList(context.Background(), &video.PublishListRequest{UserId: userID})
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
	feedResp, _ := videoClient.Feed(context.Background(), &video.FeedRequest{LatestTime: nil})
	feed := make([]*video.Video, 0)
	feed = append(feed, feedResp.VideoList)
	c.JSON(http.StatusOK, model.FeedResp{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "",
		},
		VideoList: feed,
		NextTime:  0,
	})
}

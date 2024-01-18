package favorite

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"rpc-douyin/src/config"
	"rpc-douyin/src/model"
	"rpc-douyin/src/proto/favorite"
	"rpc-douyin/src/util/connectWrapper"
	"strconv"
)

var favoriteClient favorite.FavoriteServiceClient

func init() {
	conn := connectWrapper.Connect(config.Cfg.Server.Favorite.Name)
	favoriteClient = favorite.NewFavoriteServiceClient(conn)
}

func FavoriteActionHanler(c *gin.Context) {
	userID, _ := c.Get("user_id")
	videoID, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	_, err := favoriteClient.FavoriteAction(context.Background(), &favorite.FavoriteActionRequest{
		UserId:     userID.(int64),
		VideoId:    videoID,
		ActionType: actionType,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.Response{
			StatusCode: 1,
			StatusMsg:  "点赞失败",
		})
		return
	}
	c.JSON(http.StatusOK, model.Response{
		StatusCode: 0,
		StatusMsg:  "点赞成功",
	})
}

func FavoriteListHandler(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.Query("user_id"), 10, 64)
	resp, err := favoriteClient.FavoriteList(context.Background(), &favorite.FavoriteListRequest{UserId: userID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.FavoriteListResp{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "喜欢列表获取失败",
			},
			VideoList: nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.FavoriteListResp{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "喜欢列表获取成功",
		},
		VideoList: resp.GetVideoList(),
	})
}

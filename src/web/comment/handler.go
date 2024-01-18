package comment

import (
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"rpc-douyin/src/config"
	"rpc-douyin/src/model"
	"rpc-douyin/src/proto/comment"
	"rpc-douyin/src/util/connectWrapper"
	"strconv"
)

var commentClient comment.CommentServiceClient

func init() {
	conn := connectWrapper.Connect(config.Cfg.Server.Comment.Name)
	commentClient = comment.NewCommentServiceClient(conn)
}

func CommentActionHandler(c *gin.Context) {
	actionType, _ := strconv.ParseInt(c.Query("action_type"), 10, 64)
	switch actionType {
	case 1:
		userID, _ := c.Get("user_id")
		videoID, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
		commentContent := c.Query("comment_text")
		resp, err := commentClient.CommentPost(context.Background(), &comment.CommentPostRequest{
			UserId:      userID.(int64),
			VideoId:     videoID,
			CommentText: commentContent,
		})
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.CommentActionResp{
				Response: model.Response{
					StatusCode: 1,
					StatusMsg:  "",
				},
				Comment: nil,
			})
			return
		}
		c.JSON(http.StatusOK, model.CommentActionResp{
			Response: model.Response{
				StatusCode: 0,
				StatusMsg:  "",
			},
			Comment: resp.GetComment(),
		})
	case 2:
		commentID, _ := strconv.ParseInt(c.Query("comment_id"), 10, 64)
		_, err := commentClient.CommentDel(context.Background(), &comment.CommentDelRequest{CommentId: commentID})
		if err != nil {
			c.JSON(http.StatusInternalServerError, model.CommentActionResp{
				Response: model.Response{
					StatusCode: 1,
					StatusMsg:  "删除评论失败",
				},
				Comment: nil,
			})
			return
		}
		c.JSON(http.StatusOK, model.CommentActionResp{
			Response: model.Response{
				StatusCode: 0,
				StatusMsg:  "",
			},
			Comment: nil,
		})
	}
}

func CommentListHandler(c *gin.Context) {
	videoID, _ := strconv.ParseInt(c.Query("video_id"), 10, 64)
	resp, err := commentClient.CommentList(context.Background(), &comment.CommentListRequest{VideoId: videoID})
	if err != nil {
		c.JSON(http.StatusInternalServerError, model.CommentListResp{
			Response: model.Response{
				StatusCode: 1,
				StatusMsg:  "获取评论列表失败",
			},
			CommentList: nil,
		})
		return
	}
	c.JSON(http.StatusOK, model.CommentListResp{
		Response: model.Response{
			StatusCode: 0,
			StatusMsg:  "获取评论列表成功",
		},
		CommentList: resp.GetCommentList(),
	})
}

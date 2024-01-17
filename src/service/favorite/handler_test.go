package main

import (
	"context"
	"github.com/stretchr/testify/assert"
	"rpc-douyin/src/config"
	"rpc-douyin/src/proto/favorite"
	"rpc-douyin/src/util/connectWrapper"
	"testing"
)

func TestFavoriteServiceImpl_FavoriteAction(t *testing.T) {
	conn := connectWrapper.Connect(config.Cfg.Server.Favorite.Name)
	favoriteClient := favorite.NewFavoriteServiceClient(conn)

	t.Run("喜欢", func(t *testing.T) {
		_, err := favoriteClient.FavoriteAction(context.Background(), &favorite.FavoriteActionRequest{
			UserId:     1,
			VideoId:    1,
			ActionType: 1,
		})
		assert.NoError(t, err)
	})

	t.Run("取消喜欢", func(t *testing.T) {
		_, err := favoriteClient.FavoriteAction(context.Background(), &favorite.FavoriteActionRequest{
			UserId:     1,
			VideoId:    1,
			ActionType: 2,
		})
		assert.NoError(t, err)
	})
}

func TestFavoriteServiceImpl_IsFavorite(t *testing.T) {
	conn := connectWrapper.Connect(config.Cfg.Server.Favorite.Name)
	favoriteClient := favorite.NewFavoriteServiceClient(conn)

	t.Run("是否喜欢", func(t *testing.T) {
		_, err := favoriteClient.FavoriteAction(context.Background(), &favorite.FavoriteActionRequest{
			UserId:     1,
			VideoId:    1,
			ActionType: 1,
		})
		assert.NoError(t, err)

		isFavor, err := favoriteClient.IsFavorite(context.Background(), &favorite.IsFavoriteRequest{
			UserId:  1,
			VideoId: 1,
		})
		assert.NoError(t, err)
		assert.True(t, isFavor.GetFavorite())

		_, err = favoriteClient.FavoriteAction(context.Background(), &favorite.FavoriteActionRequest{
			UserId:     1,
			VideoId:    1,
			ActionType: 2,
		})
		assert.NoError(t, err)

		isFavor, err = favoriteClient.IsFavorite(context.Background(), &favorite.IsFavoriteRequest{
			UserId:  1,
			VideoId: 1,
		})
		assert.NoError(t, err)
		assert.False(t, isFavor.GetFavorite())
	})
}

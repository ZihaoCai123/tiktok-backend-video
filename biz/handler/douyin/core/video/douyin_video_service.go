// Code generated by hertz generator.

package video

import (
	"BiteDans.com/tiktok-backend/biz/dal"
	"BiteDans.com/tiktok-backend/biz/dal/model"
	"BiteDans.com/tiktok-backend/biz/model/douyin/core/user"
	"BiteDans.com/tiktok-backend/biz/model/douyin/core/video"
	"BiteDans.com/tiktok-backend/pkg/utils"
	"context"
	"fmt"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// VideoFeed .
// @router /douyin/feed [GET]
func VideoFeed(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.DouyinVideoFeedRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video.DouyinVideoFeedResponse)

	c.JSON(consts.StatusOK, resp)
}

// VideoPublish .
// @router /douyin/publish/action/ [POST]
func VideoPublish(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.DouyinVideoPublishRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video.DouyinVideoPublishResponse)

	c.JSON(consts.StatusOK, resp)
}

// VideoPublishList .
// @router /douyin/publish/list/ [GET]
func VideoPublishList(ctx context.Context, c *app.RequestContext) {
	var err error
	var req video.DouyinVideoPublishListRequest

	err = c.BindAndValidate(&req)
	fmt.Println(req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(video.DouyinVideoPublishListResponse)

	if _, err = utils.GetIdFromToken(req.Token); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "Invalid token"
		resp.VideoList = nil

		c.JSON(consts.StatusUnauthorized, resp)
		return
	}

	_user := new(model.User)

	if err = model.FindUserById(_user, uint(req.UserId)); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "User id does not exist"
		resp.VideoList = nil
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	var _videos []*model.Video

	if err = model.FindVideosByUserId(_videos, int64(_user.ID)); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "Fail to retrieve videos from User id"
		resp.VideoList = nil
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	dal.DB.Where("author_id = ?", int64(_user.ID)).Find(&_videos)

	resp.StatusCode = 0
	resp.StatusMsg = "Publishing list info retrieved successfully"
	resp.VideoList = []*video.Video{}

	for _, _video := range _videos {
		the_user := &user.User{
			ID:            int64(_user.ID),
			Name:          _user.Username,
			FollowCount:   123,
			FollowerCount: 456,
			IsFollow:      true,
		}
		the_video := &video.Video{
			ID:            int64(_video.ID),
			Author:        (*video.User)(the_user),
			PlayUrl:       _video.PlayUrl,
			CoverUrl:      _video.CoverUrl,
			FavoriteCount: _video.FavoriteCount,
			CommentCount:  _video.CommentCount,
			IsFavorite:    false,
			Title:         _video.Title,
		}
		resp.VideoList = append(resp.VideoList, the_video)
	}

	c.JSON(consts.StatusOK, resp)
}
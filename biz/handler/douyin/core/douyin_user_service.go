// Code generated by hertz generator.

package core

import (
	"context"

	"BiteDans.com/tiktok-backend/biz/dal/model"
	core "BiteDans.com/tiktok-backend/biz/model/douyin/core"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
)

// UserInfo .
// @router /douyin/user [GET]
func UserInfo(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinUserRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinUserResponse)

	if _error := c.Errors.Last(); _error != nil {
		resp.StatusCode = -1
		resp.StatusMsg = _error.Error()
		resp.User = nil

		c.JSON(consts.StatusUnauthorized, resp)
		return
	}

	user := new(model.User)

	if err = model.FindUserById(user, uint(req.UserId)); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "User id does not exist"
		resp.User = nil
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "User info retrieved successfully"
	resp.User = &core.User{
		ID:            int64(user.ID),
		Name:          user.Username,
		FollowCount:   123,
		FollowerCount: 456,
		IsFollow:      true,
	}

	c.JSON(consts.StatusOK, resp)
}

// UserLogin .
// @router /douyin/usr/login [POST]
func UserLogin(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinUserLoginRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinUserLoginResponse)
	user := new(model.User)

	user.Username = req.Username
	user.Password = req.Password

	var inputpassword = user.Password

	if err = model.FindUserByUsername(user, req.Username); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "Failed to log in (Username not found)"
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	if inputpassword != user.Password {
		resp.StatusCode = -1
		resp.StatusMsg = "Failed to log in (Incorrect password)"
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "User logged in successfully"
	resp.UserId = int64(user.ID)
	resp.Token = "token"

	c.JSON(consts.StatusOK, resp)
}

// UserRegister .
// @router /douyin/user/register/ [POST]
func UserRegister(ctx context.Context, c *app.RequestContext) {
	var err error
	var req core.DouyinUserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		c.String(consts.StatusBadRequest, err.Error())
		return
	}

	resp := new(core.DouyinUserRegisterResponse)
	user := new(model.User)

	if err = model.FindUserByUsername(user, req.Username); err == nil {
		resp.StatusCode = -1
		resp.StatusMsg = "Username has been used"
		c.JSON(consts.StatusBadRequest, resp)
		return
	}

	user.Username = req.Username
	user.Password = req.Password

	if err = model.CreateUser(user); err != nil {
		resp.StatusCode = -1
		resp.StatusMsg = "Failed to register user"
		c.JSON(consts.StatusInternalServerError, resp)
		return
	}

	resp.StatusCode = 0
	resp.StatusMsg = "User registered successfully"
	resp.UserId = int64(user.ID)
	resp.Token = "token"

	c.JSON(consts.StatusOK, resp)
}

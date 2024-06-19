package controllers

import (
	"dxj/internal/pkg/e"
	"dxj/internal/pkg/requests"
	"dxj/internal/services"

	"github.com/gin-gonic/gin"
)

type PublicController struct {
}

var userService = services.UserService{}

func (PublicController) GetHello(c *gin.Context) {
	e.SendData(c, e.SUCCESS, "hello world")
}

// @Summary 用户登录
// @Description 通过用户名和密码登录系统
// @Tags 认证
// @Accept json
// @Produce json
// @Param loginUser body requests.LoginUserReq true "登录请求体"
// @Success 200 {object} response.LoginUserResp "登录成功，返回用户信息及token"
// @Failure 400 {string} string "请求参数无效"
// @Failure 401 {string} string "用户名或密码错误"
// @Router /public/login [post]
func (PublicController) Login(c *gin.Context) {
	var loginUser *requests.LoginUserReq
	err := c.ShouldBindJSON(&loginUser)
	if err != nil {
		e.Send(c, e.INVALID_PARAMS)
		return
	}
	loginUserResp := userService.Login(loginUser)
	if loginUserResp != nil {
		e.SendData(c, e.SUCCESS, loginUserResp)

	} else {
		e.SendData(c, e.ERROR, "用户名或密码错误")
		return
	}
}

// 用户注册
func (PublicController) Register(c *gin.Context) {
	var registerUser *requests.RegisterUserReq
	err := c.ShouldBindJSON(&registerUser)
	if err != nil {
		e.Send(c, e.INVALID_PARAMS)
		return
	}
	user := userService.Register(registerUser)
	if user != nil {
		e.SendData(c, e.SUCCESS, user)
	} else {
		e.Send(c, e.ERROR)
	}
}

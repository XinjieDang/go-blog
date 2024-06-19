package controllers

import (
	"dxj/internal/pkg/e"
	"dxj/internal/services"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
}

var userServiceImpl = services.UserService{}

// @Summary 测试接口
// @Description 该接口用于获取当前登录用户的信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {string} string "成功返回用户信息，例如: 'John Doe'"
// @Router /user/getHello [get]
func (UserController) GetUserInfo(c *gin.Context) {
	e.SendData(c, e.SUCCESS, "ok")
}

// @Summary 通过id获取用户信息
// @Description 该接口用于获取当前登录用户的信息
// @Tags 用户管理
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.SuccessResponse "成功返回用户信息"
// @Router /user/getUserInfoById/:id [get]
func (UserController) GetUserInfoById(c *gin.Context) {
	id := c.Param("id")
	intNumber, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		fmt.Printf("转换错误: %v\n", err)
		return
	}
	resp := userServiceImpl.GetUserInfo(intNumber)
	e.SendData(c, e.SUCCESS, resp)
}

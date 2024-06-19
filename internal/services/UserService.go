package services

import (
	"dxj/internal/global"
	"dxj/internal/models"
	"dxj/internal/pkg/requests"
	"dxj/internal/pkg/response"
	"dxj/internal/pkg/utils"
	"dxj/internal/repository/impl"
	"time"
)

/**
* @Description:用户登录
**/
type UserService struct {
}

var userRepo = impl.UserRepoImpl{}

// 用户登录，颁发token
func (userService *UserService) Login(loginUser *requests.LoginUserReq) *response.LoginUserResp {
	//查询用户是否存在
	user := models.User{}
	result := global.DB.Where("user_name = ?", loginUser.UserName).First(&user)
	if result.Error != nil {
		return nil
	}
	//验证密码
	if user.PassWord != loginUser.PassWord {
		return nil
	}
	//生成token
	jwtConfig := global.Config.Jwt.User
	//fmt.Println("xinfo", global.Config.Jwt.Admin)
	token, err := utils.GenerateToken(uint64(user.ID), jwtConfig.Name, jwtConfig.Secret)
	if err != nil {
		return nil
	}
	return &response.LoginUserResp{UserName: user.UserName, Avatar: user.Avatar, Token: token}
}

// 用户注册
func (userService *UserService) Register(registerUser *requests.RegisterUserReq) *models.User {
	user := models.User{}
	user.UserName = registerUser.UserName
	user.PassWord = registerUser.PassWord
	user.Email = registerUser.Email
	user.Phone = registerUser.Phone
	user.Bio = registerUser.Bio
	user.Avatar = registerUser.Avatar
	//上次登录的时间为当前时间
	user.LastLoginTime = time.Now()
	//user.LastLoginTime=time.Date()
	global.DB.Create(&user)
	return &user
}

// 获取用户信息
func (userService *UserService) GetUserInfo(id int64) *response.UserInfoResp {
	return userRepo.GetUserById(id)
}

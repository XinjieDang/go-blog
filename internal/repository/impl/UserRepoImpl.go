package impl

import (
	"dxj/internal/global"
	"dxj/internal/pkg/response"
)

type UserRepoImpl struct {
}

func (UserRepo *UserRepoImpl) GetUserById(id int64) *response.UserInfoResp {
	UserInfoResp := response.UserInfoResp{}
	global.DB.Table("users").Select("user_name", "email", "phone", "avatar", "status", "bio", "is_active", "last_login_time").Where("id = ?", id).First(&UserInfoResp)
	return &UserInfoResp
}

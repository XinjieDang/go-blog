package response

import "time"

type UserInfoResp struct {
	UserName      string    `json:"user_name"`
	Email         string    `json:"email"`
	Phone         string    `json:"phone"`
	Avatar        string    `json:"avatar"`
	Status        int64     `json:"status" gorm:"default:0"`
	Bio           string    `json:"bio"`                        //简介
	IsActive      int64     `json:"is_active" gorm:"default:0"` //是否激活
	LastLoginTime time.Time `json:"last_login_time"`            //上次登录时间
}
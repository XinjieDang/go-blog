package response

import "time"

type LoginUserResp struct {
	UserName      string    `json:"user_name"`
	Avatar        string    `json:"avatar"`
	LastLoginTime time.Time `json:"last_login_time"`
	Token         string    `json:"token"`
}

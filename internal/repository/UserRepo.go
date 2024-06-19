package repository

import "dxj/internal/pkg/response"

type UserRepo interface {
	GetUserById(id int64) *response.UserInfoResp
}

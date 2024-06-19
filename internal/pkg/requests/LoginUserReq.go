package requests

type LoginUserReq struct {
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
}

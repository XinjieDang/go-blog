package requests

type RegisterUserReq struct {
	UserName string `json:"user_name"`
	PassWord string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Bio      string `json:"bio"`
	Avatar   string `json:"avatar"`
}

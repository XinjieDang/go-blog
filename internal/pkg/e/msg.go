package e

import (
	"dxj/internal/pkg/response"
	"net/http"

	"github.com/gin-gonic/gin"
)

var ErrMsg = map[int]string{
	SUCCESS:         "ok",
	ERROR:           "内部错误",
	UNKNOW_IDENTITY: "未知身份",
	INVALID_PARAMS:  "无效参数",
}

func GetMsg(code int) string {
	return ErrMsg[code]
}

func Send(ctx *gin.Context, code int) {
	ctx.JSON(http.StatusOK, response.Result{Code: code, Msg: GetMsg(code)})
}

func SendData(ctx *gin.Context, code int, data interface{}) {
	ctx.JSON(http.StatusOK, response.Result{Code: code, Msg: GetMsg(code), Data: data})
}

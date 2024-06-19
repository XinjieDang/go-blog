package controllers

import (
	"dxj/internal/global"
	"dxj/internal/pkg/e"
	"dxj/internal/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"os"
	"path/filepath"
	"time"
)

type FileUploadController struct {
}

// @Summary 通用文件上传-单文件
// @Description 该接口用于获处理文件上传
// @Tags 文件上传
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} response.SuccessResponse "成功返回文件信息，例如: 'xxxxx'"
// @Router /public/uploadFile [post]
func (FileUploadController) UploadFile(c *gin.Context) {
	logger := logger.LogrusLogger
	// 单文件上传
	file, err := c.FormFile("file")
	if err != nil {
		e.SendData(c, e.INVALID_PARAMS, err.Error())
		return
	}
	// 指定保存文件的目录为D盘的upload
	saveDir := global.Config.Upload.Path
	if err := os.MkdirAll(saveDir, os.ModePerm); err != nil {
		e.SendData(c, e.ERROR, err.Error())
		return
	}
	// 生成新的文件名，例如加上时间戳避免同名冲突
	ext := filepath.Ext(file.Filename)
	newFileName := fmt.Sprintf("%d%s", time.Now().UnixNano(), ext)
	// 构建完整的文件保存路径
	dst := filepath.Join(saveDir, newFileName)
	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		e.SendData(c, e.ERROR, err.Error())
		return
	}
	defer src.Close()
	// 创建目标文件
	out, err := os.Create(dst)
	if err != nil {
		e.SendData(c, e.ERROR, err.Error())
		return
	}
	defer out.Close()
	// 复制文件内容
	_, err = io.Copy(out, src)
	if err != nil {
		e.SendData(c, e.ERROR, err.Error())
		return
	}
	logger.Printf(`🍟: upload file Successfully  ` + newFileName)
	e.SendData(c, e.SUCCESS, newFileName)
}

package controllers

import (
	"dxj/internal/pkg/e"
	"dxj/internal/pkg/requests"
	"dxj/internal/services"

	"github.com/gin-gonic/gin"
)

var service = services.CategoryService{}

type CategoryController struct {
	CategoryService *services.CategoryService
}

func NewCategoryController() *CategoryController {
	return &CategoryController{
		CategoryService: services.NewCategoryService(),
	}
}

// 创建分类
func (controller *CategoryController) CreateCategory(c *gin.Context) {
	var data map[string]interface{}
	if err := c.ShouldBindJSON(&data); err != nil {
		e.SendData(c, e.INVALID_PARAMS, err)
		return
	}
	category := service.CreateCategory(data)
	e.SendData(c, e.SUCCESS, category)
}

// 分页获取分类列表
func (controller *CategoryController) GetCategoryList(c *gin.Context) {
	var req requests.CategoryListReq
	if err := c.BindJSON(&req); err != nil {
		e.SendData(c, e.INVALID_PARAMS, err)
		return
	}
	categories := service.GetCategoryList(&req)
	e.SendData(c, e.SUCCESS, categories)
}

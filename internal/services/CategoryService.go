package services

/*
 * @LastEditTime: 2021-04-06 15:06:08
 * @LastEditors: dxj
 * @Description: 分类服务
 */
import (
	"dxj/internal/global"
	"dxj/internal/models"
	"dxj/internal/pkg/dto"
	"dxj/internal/pkg/requests"
	"dxj/internal/pkg/utils"
)

type CategoryService struct {
}

func NewCategoryService() *CategoryService {
	return &CategoryService{}
}

// 分页获取分类列表
func (categoryService *CategoryService) GetCategoryList(req *requests.CategoryListReq) []*models.Category {
	var categoryList []*models.Category
	query := global.DB
	if req != nil {
		if req.Name != "" {
			query.Where("name like ?", "%"+req.Name+"%")
		}
	}
	pageReq := &dto.PageParam{
		Page:     req.Page,
		PageSize: req.PageSize,
	}
	query.Scopes(utils.Paginate(pageReq)).Find(&categoryList)
	return categoryList
}

// 创建分类
func (categoryService *CategoryService) CreateCategory(data map[string]interface{}) *models.Category {
	var category models.Category
	category.Name = data["name"].(string)
	category.Desc = data["desc"].(string)
	//category.IsActive = data["is_active"].(int64)
	//category.SortOrder = data["sort_order"].(int64)
	global.DB.Create(&category)
	return &category
}

// 删除分类
func (categoryService *CategoryService) DeleteCategory(id int64) error {
	var category models.Category
	global.DB.Where("id = ?", id).Delete(&category)
	return nil
}

// 通过id获取分类
func (categoryService *CategoryService) GetCategoryById(id int64) *models.Category {
	var category models.Category
	global.DB.Where("id = ?", id).First(&category)
	return &category
}

// 更新分类
func (categoryService *CategoryService) UpdateCategory(id int64, data map[string]interface{}) *models.Category {
	var category models.Category
	global.DB.Where("id = ?", id).First(&category)
	category.Name = data["name"].(string)
	category.Desc = data["desc"].(string)
	//category.IsActive = data["is_active"].(int64)
	//category.SortOrder= data["sort_order"].(int64)
	global.DB.Save(&category)
	return &category
}

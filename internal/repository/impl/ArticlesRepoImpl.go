package impl

import (
	"dxj/internal/global"
	"dxj/internal/models"
	"dxj/internal/pkg/response"
)

type ArticlesRepoImpl struct {
}

// 新建文章
func (articleRepo *ArticlesRepoImpl) CreateArticle(article *models.Article) error {
	return global.DB.Create(&article).Error
}

// 通过文章id获取文章详情
func (articleRepo *ArticlesRepoImpl) GetArticleById(id int64) *response.ArticleDetailResp {
	var result response.ArticleDetailResp
	global.DB.Table("articles").
		Select("articles.title, users.user_name as author, categories.name as category").
		Joins("JOIN users ON articles.author_id = users.id").
		Joins("JOIN categories ON articles.category_id = categories.id").
		Where("articles.id = ?", id).
		Scan(&result)
	return &result
}

// 获取文章详情列表分页
func (articleRepo *ArticlesRepoImpl) GetArticleList(page, pageSize int64) ([]response.ArticleDetailResp, int64) {
	var result []response.ArticleDetailResp
	var count int64
	global.DB.Table("articles").
		Select("articles.title, users.user_name as author, categories.name as category").
		Joins("JOIN users ON articles.author_id = users.id").
		Joins("JOIN categories ON articles.category_id = categories.id")
	if page > 0 && pageSize > 0 {
		global.DB.Offset(int((page - 1) * pageSize)).Limit(int(pageSize)).Scan(&result)
		global.DB.Table("articles").Count(&count)
	} else {
		global.DB.Scan(&result)
		global.DB.Table("articles").Count(&count)
		return result, count
	}
	return result, count
}

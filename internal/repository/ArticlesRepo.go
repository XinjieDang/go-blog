package repository

import (
	"dxj/internal/models"
	"dxj/internal/pkg/response"
)

type ArticlesRepo interface {
	//新建文章
	CreateArticle(article *models.Article) error
	//通过id查询文章详情
	GetArticleById(id int64) *response.ArticleDetailResp
	//获取文章列表
	GetArticleList(page, pageSize int64) ([]response.ArticleDetailResp, int64)
}

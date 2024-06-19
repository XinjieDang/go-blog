package services

import (
	"dxj/internal/models"
	"dxj/internal/pkg/response"
	"dxj/internal/repository/impl"
)

type ArticleService struct {
}

var articlesRepo = impl.ArticlesRepoImpl{}

// 新建文章
func (articleService *ArticleService) CreateArticle(data *models.Article) *models.Article {
	article := &models.Article{
		Content:      data.Content,
		IsTop:        data.IsTop,
		AuthorId:     data.AuthorId,
		CategoryId:   data.CategoryId,
		LikeCount:    data.LikeCount,
		CommentCount: data.CommentCount,
		ViewCount:    data.ViewCount,
		Title:        data.Title,
		Status:       data.Status,
		IsHot:        data.IsHot,
	}
	articlesRepo.CreateArticle(article)
	return article
}

// 通过文章id获取文章详情
func (articleService *ArticleService) GetArticleById(id int64) *response.ArticleDetailResp {
	return articlesRepo.GetArticleById(id)
}

// 分页获取文章详情列表
func (articleService *ArticleService) GetArticleList(page, pageSize int64) ([]response.ArticleDetailResp, int64) {
	return articlesRepo.GetArticleList(page, pageSize)
}

package repository

import "dxj/internal/models"

type CategoryRepo interface {
	//获取分类信息
	GetCategory(id int) (*models.Category, error)
}

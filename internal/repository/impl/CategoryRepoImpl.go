package impl

import (
	"dxj/internal/models"

	"gorm.io/gorm"
)

type CategoryRepoImpl struct {
	db *gorm.DB
}

func NewCategoryRepoImpl(db *gorm.DB) *CategoryRepoImpl {
	return &CategoryRepoImpl{
		db: db,
	}
}

func (CategoryRepo *CategoryRepoImpl) GetCategory(id int) (*models.Category, error) {
	var Category models.Category
	CategoryRepo.db.Where("id = ?", id).First(&Category)
	return &Category, nil
}

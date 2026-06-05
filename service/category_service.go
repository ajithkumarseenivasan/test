package service

import (
	"time"
	"user-management/model"
	"user-management/repository"
)

type CategoryService interface {
	SaveCategory(categoryRequest model.CategoryRequest) (bool, error)
	GetCategories(tenantId string, page int64, limit int64) ([]model.Category, int64, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) CategoryService {
	return &categoryService{repo: r}
}

func (c *categoryService) SaveCategory(categoryRequest model.CategoryRequest) (bool, error) {
	categoryRequest.Category.TenantID = categoryRequest.TenantId
	categoryRequest.Category.CreatedBy = categoryRequest.UserId
	categoryRequest.Category.ModifiedDate = time.Now().UTC()
	categoryRequest.Category.CreatedDate = time.Now().UTC()
	return c.repo.SaveCategory(categoryRequest.Category)
}

func (c *categoryService) GetCategories(tenantId string, page int64, limit int64) ([]model.Category, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	skip := (page - 1) * limit
	return c.repo.GetCategories(tenantId, skip, limit)
}

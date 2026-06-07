package service

import (
	"time"
	"user-management/model"
	"user-management/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type CategoryService interface {
	SaveCategory(categoryRequest model.CategoryRequest) (bool, error)
	GetCategories(tenantId string, name string, description string, page int64, limit int64) ([]model.Category, int64, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) CategoryService {
	return &categoryService{repo: r}
}

func (c *categoryService) SaveCategory(categoryRequest model.CategoryRequest) (bool, error) {
	categoryRequest.Category.TenantID, _ = primitive.ObjectIDFromHex(categoryRequest.TenantId)
	categoryRequest.Category.CreatedBy, _ = primitive.ObjectIDFromHex(categoryRequest.UserId)
	categoryRequest.Category.ModifiedDate = time.Now().UTC()
	categoryRequest.Category.CreatedDate = time.Now().UTC()
	return c.repo.SaveCategory(categoryRequest.Category)
}

func (c *categoryService) GetCategories(tenantId string, name string, description string, page int64, limit int64) ([]model.Category, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	skip := (page - 1) * limit
	return c.repo.GetCategories(tenantId, name, description, skip, limit)
}

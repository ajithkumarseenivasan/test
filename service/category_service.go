package service

import (
	"user-management/model"
	"user-management/repository"
)

type CategoryService interface {
	SaveCategory(category model.Category) (bool, error)
}

type categoryService struct {
	repo repository.CategoryRepository
}

func NewCategoryService(r repository.CategoryRepository) CategoryService {
	return &categoryService{repo: r}
}

func (c *categoryService) SaveCategory(category model.Category) (bool, error) {
	return c.repo.SaveCategory(category)
}

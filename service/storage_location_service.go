package service

import (
	"time"
	"user-management/model"
	"user-management/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type StorageLocationService interface {
	SaveLocation(req model.StorageLocationRequest) (bool, error)
	GetLocations(tenantId string, name string, description string, page int64, limit int64) ([]model.StorageLocation, int64, error)
}

type storageLocationService struct {
	repo repository.StorageLocationRepository
}

func NewStorageLocationService(r repository.StorageLocationRepository) StorageLocationService {
	return &storageLocationService{repo: r}
}

func (s *storageLocationService) SaveLocation(req model.StorageLocationRequest) (bool, error) {
	req.Location.TenantID, _ = primitive.ObjectIDFromHex(req.TenantId)
	req.Location.CreatedBy, _ = primitive.ObjectIDFromHex(req.UserId)
	req.Location.ModifiedDate = time.Now().UTC()
	req.Location.CreatedDate = time.Now().UTC()
	return s.repo.SaveLocation(req.Location)
}

func (s *storageLocationService) GetLocations(tenantId string, name string, description string, page int64, limit int64) ([]model.StorageLocation, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	skip := (page - 1) * limit
	return s.repo.GetLocations(tenantId, name, description, skip, limit)
}

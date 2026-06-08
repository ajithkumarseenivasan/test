package service

import (
	"time"
	"user-management/model"
	"user-management/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InPlantUnitService interface {
	SaveInPlantUnit(inPlantUnitRequest model.InPlantUnitRequest) (bool, error)
	GetInPlantUnits(tenantId string, name string, description string, page int64, limit int64) ([]model.InPlantUnit, int64, error)
}

type inPlantUnitService struct {
	repo repository.InPlantUnitRepository
}

func NewInPlantUnitService(r repository.InPlantUnitRepository) InPlantUnitService {
	return &inPlantUnitService{repo: r}
}

func (i *inPlantUnitService) SaveInPlantUnit(inPlantUnitRequest model.InPlantUnitRequest) (bool, error) {
	inPlantUnitRequest.InPlantUnit.TenantID, _ = primitive.ObjectIDFromHex(inPlantUnitRequest.TenantId)
	inPlantUnitRequest.InPlantUnit.CreatedBy, _ = primitive.ObjectIDFromHex(inPlantUnitRequest.UserId)
	inPlantUnitRequest.InPlantUnit.ModifiedDate = time.Now().UTC()
	inPlantUnitRequest.InPlantUnit.CreatedDate = time.Now().UTC()
	return i.repo.SaveInPlantUnit(inPlantUnitRequest.InPlantUnit)
}

func (i *inPlantUnitService) GetInPlantUnits(tenantId string, name string, description string, page int64, limit int64) ([]model.InPlantUnit, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	skip := (page - 1) * limit
	return i.repo.GetInPlantUnits(tenantId, name, description, skip, limit)
}

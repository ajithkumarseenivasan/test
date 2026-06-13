package service

import (
	"time"
	"user-management/model"
	"user-management/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryItemService interface {
	SaveInventoryItem(itemRequest model.InventoryItemRequest) (bool, error)
	GetInventoryItems(tenantId string, name string, vendorId string, page int64, limit int64) ([]model.InventoryItemUI, int64, error)
}

type inventoryItemService struct {
	repo repository.InventoryItemRepository
}

func NewInventoryItemService(r repository.InventoryItemRepository) InventoryItemService {
	return &inventoryItemService{repo: r}
}

func (s *inventoryItemService) SaveInventoryItem(itemRequest model.InventoryItemRequest) (bool, error) {
	itemRequest.InventoryItem.TenantID, _ = primitive.ObjectIDFromHex(itemRequest.TenantId)
	itemRequest.InventoryItem.CreatedBy, _ = primitive.ObjectIDFromHex(itemRequest.UserId)
	itemRequest.InventoryItem.ModifiedDate = time.Now().UTC()
	itemRequest.InventoryItem.CreatedDate = time.Now().UTC()
	return s.repo.SaveInventoryItem(itemRequest.InventoryItem)
}

func (s *inventoryItemService) GetInventoryItems(tenantId string, name string, vendorId string, page int64, limit int64) ([]model.InventoryItemUI, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	skip := (page - 1) * limit
	return s.repo.GetInventoryItems(tenantId, name, vendorId, skip, limit)
}

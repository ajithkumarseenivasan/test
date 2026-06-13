package service

import (
	"time"
	"user-management/model"
	"user-management/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PurchaseOrderService interface {
	SavePurchaseOrder(itemRequest model.PurchaseOrderRequest) (bool, error)
	GetPurchaseOrders(tenantId, status, vendorName, orderId string, page int64, limit int64) ([]model.PurchaseOrder, int64, error)
}

type purchaseOrderService struct {
	repo repository.PurchaseOrderRepository
}

func NewPurchaseOrderService(r repository.PurchaseOrderRepository) PurchaseOrderService {
	return &purchaseOrderService{repo: r}
}

func (s *purchaseOrderService) SavePurchaseOrder(orderRequest model.PurchaseOrderRequest) (bool, error) {
	orderRequest.PurchaseOrder.TenantID, _ = primitive.ObjectIDFromHex(orderRequest.TenantId)
	orderRequest.PurchaseOrder.CreatedBy, _ = primitive.ObjectIDFromHex(orderRequest.UserId)
	orderRequest.PurchaseOrder.ModifiedDate = time.Now().UTC()
	orderRequest.PurchaseOrder.CreatedDate = time.Now().UTC()
	orderRequest.PurchaseOrder.Status = "Placed"
	return s.repo.SavePurchaseOrder(orderRequest.PurchaseOrder)
}

func (s *purchaseOrderService) GetPurchaseOrders(tenantId, status, vendorName, orderId string, page int64, limit int64) ([]model.PurchaseOrder, int64, error) {
	if page < 1 {
		page = 1
	}
	if limit < 1 {
		limit = 10
	}
	skip := (page - 1) * limit
	return s.repo.GetPurchaseOrders(tenantId, status, vendorName, orderId, skip, limit)
}

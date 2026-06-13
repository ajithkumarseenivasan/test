package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type PurchaseOrder struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	TenantID     primitive.ObjectID `bson:"tenantId" json:"tenantId"`
	VendorId     primitive.ObjectID `bson:"vendorId" json:"vendorId"`
	VendorName   string             `bson:"vendorName" json:"vendorName"`
	OrderDate    string             `bson:"orderDate" json:"orderDate"`
	OrderItems   []OrderItem        `bson:"orderItems" json:"orderItems"`
	CreatedDate  time.Time          `bson:"createdDate" json:"createdDate"`
	ModifiedDate time.Time          `bson:"modifiedDate" json:"modifiedDate"`
	CreatedBy    primitive.ObjectID `bson:"createdBy" json:"createdBy"`
	Status       string             `bson:"status" json:"status"`
}

type OrderItem struct {
	ItemId      primitive.ObjectID `bson:"itemId,omitempty" json:"itemId,omitempty"`
	ItemName    string             `bson:"itemName" json:"itemName"`
	InPlantUnit string             `bson:"inPlantUnit" json:"inPlantUnit"`
	UnitPrice   float64            `bson:"unitPrice" json:"unitPrice"`
	Price       float64            `bson:"price" json:"price"`
	Quantity    float64            `bson:"quantity" json:"quantity"`
}

type PurchaseOrderListResponse struct {
	PurchaseOrders []PurchaseOrder `json:"data"`
	Page           int64           `json:"page"`
	Limit          int64           `json:"limit"`
	Total          int64           `json:"total"`
}

type PurchaseOrderRequest struct {
	MasterUiRequest
	PurchaseOrder PurchaseOrder `json:"purchaseOrder"`
}

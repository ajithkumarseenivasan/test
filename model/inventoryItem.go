package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InventoryItem struct {
	ID                primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name              string             `bson:"name" json:"name"`
	InventoryUnitId   primitive.ObjectID `bson:"inventoryUnitId" json:"inventoryUnitId"`
	CategoryId        primitive.ObjectID `bson:"categoryId" json:"categoryId"`
	StorageLocationId primitive.ObjectID `bson:"storageLocationId" json:"storageLocationId"`
	ParLevel          int                `bson:"parLevel" json:"parLevel"`
	Vendors           []InventoryVendor  `bson:"vendors" json:"vendors"`
	TenantID          primitive.ObjectID `bson:"tenantId" json:"tenantId"`
	CreatedDate       time.Time          `bson:"createdDate" json:"createdDate"`
	ModifiedDate      time.Time          `bson:"modifiedDate" json:"modifiedDate"`
	CreatedBy         primitive.ObjectID `bson:"createdBy" json:"createdBy"`
}

type InventoryItemUI struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `bson:"name" json:"name"`
	Category     string             `bson:"category" json:"category"`
	Storage      string             `bson:"storage" json:"storage"`
	TenantID     primitive.ObjectID `bson:"tenantId" json:"tenantId"`
	ParLevel     int                `bson:"parLevel" json:"parLevel"`
	Vendors      []InventoryVendor  `bson:"vendors" json:"vendors"`
	Unit         string             `bson:"unit" json:"unit"`
	CreatedDate  time.Time          `bson:"createdDate" json:"createdDate"`
	ModifiedDate time.Time          `bson:"modifiedDate" json:"modifiedDate"`
	CreatedBy    primitive.ObjectID `bson:"createdBy" json:"createdBy"`
}

type InventoryVendor struct {
	ID    primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Price float64            `bson:"price" json:"price"`
}

type InventoryItemListResponse struct {
	InventoryItems []InventoryItemUI `json:"data"`
	Page           int64             `json:"page"`
	Limit          int64             `json:"limit"`
	Total          int64             `json:"total"`
}

type InventoryItemRequest struct {
	MasterUiRequest
	InventoryItem InventoryItem `json:"inventoryItem"`
}

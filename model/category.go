package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Category struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `bson:"name" json:"name"`
	Description  string             `bson:"description" json:"description"`
	Code         string             `bson:"code" json:"code"`
	TenantID     string             `bson:"tenantId" json:"tenantId"`
	CreatedDate  time.Time          `bson:"createdDate" json:"createdDate"`
	ModifiedDate time.Time          `bson:"modifiedDate" json:"modifiedDate"`
	CreatedBy    string             `bson:"createdBy" json:"createdBy"`
}

type CategoryListResponse struct {
	Categories []Category `json:"data"`
	Page       int64      `json:"page"`
	Limit      int64      `json:"limit"`
	Total      int64      `json:"total"`
}

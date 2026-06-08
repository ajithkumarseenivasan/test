package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type InPlantUnit struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name         string             `bson:"name" json:"name"`
	Description  string             `bson:"description" json:"description"`
	Code         string             `bson:"code" json:"code"`
	TenantID     primitive.ObjectID `bson:"tenantId" json:"tenantId"`
	CreatedDate  time.Time          `bson:"createdDate" json:"createdDate"`
	ModifiedDate time.Time          `bson:"modifiedDate" json:"modifiedDate"`
	CreatedBy    primitive.ObjectID `bson:"createdBy" json:"createdBy"`
}

type InPlantUnitListResponse struct {
	InPlantUnits []InPlantUnit `json:"data"`
	Page         int64         `json:"page"`
	Limit        int64         `json:"limit"`
	Total        int64         `json:"total"`
}

type InPlantUnitRequest struct {
	MasterUiRequest
	InPlantUnit InPlantUnit `json:"inPlantUnit"`
}

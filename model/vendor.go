package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Vendor struct {
	ID                 primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name               string             `bson:"name" json:"name"`
	Description        string             `bson:"description" json:"description"`
	Email              string             `bson:"email" json:"email"`
	TenantID           primitive.ObjectID `bson:"tenantId" json:"tenantId"`
	CreatedDate        time.Time          `bson:"createdDate" json:"createdDate"`
	ModifiedDate       time.Time          `bson:"modifiedDate" json:"modifiedDate"`
	CreatedBy          primitive.ObjectID `bson:"createdBy" json:"createdBy"`
	VendorAddress      VendorAddress      `bson:"address" json:"address"`
	ContactPersonName  string             `bson:"contactPersonName" json:"contactPersonName"`
	ContactPersonPhone string             `bson:"contactPersonPhoneNumber" json:"contactPersonPhoneNumber"`
}

type VendorAddress struct {
	AddressLine string `bson:"addressLine" json:"addressLine"`
	City        string `bson:"city" json:"city"`
	State       string `bson:"state" json:"state"`
	Country     string `bson:"country" json:"country"`
	ZipCode     string `bson:"zipCode" json:"zipCode"`
}

type VendorListResponse struct {
	Vendors []Vendor `json:"data"`
	Page    int64    `json:"page"`
	Limit   int64    `json:"limit"`
	Total   int64    `json:"total"`
}

type VendorRequest struct {
	MasterUiRequest
	Vendor Vendor `json:"vendor"`
}

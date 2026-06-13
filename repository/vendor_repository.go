package repository

import (
	"context"
	"errors"
	"time"
	"user-management/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type VendorRepository interface {
	SaveVendor(vendor model.Vendor) (bool, error)
	GetVendors(tenantId string, name string, description string, email string, contactPersonName string, skip int64, limit int64) ([]model.Vendor, int64, error)
}

type vendorRepository struct {
	vendorCollection *mongo.Collection
}

func NewVendorRepository(client *mongo.Client) VendorRepository {
	collection := client.Database("stratos").Collection("inventoryVendors")
	return &vendorRepository{vendorCollection: collection}
}

func (s *vendorRepository) SaveVendor(vendor model.Vendor) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check for duplicate name within the same tenant
	filter := bson.M{
		"name":     vendor.Name,
		"tenantId": vendor.TenantID,
	}
	count, err := s.vendorCollection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return false, errors.New("Vendor with the same name already exists.")
	}

	if vendor.ID.IsZero() {
		vendor.ID = primitive.NewObjectID()
	}
	if vendor.CreatedDate.IsZero() {
		vendor.CreatedDate = time.Now().UTC()
	}
	vendor.ModifiedDate = time.Now().UTC()

	_, err = s.vendorCollection.InsertOne(ctx, vendor)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *vendorRepository) GetVendors(tenantId string, name string, description string, email string, contactPersonName string, skip int64, limit int64) ([]model.Vendor, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	if tenantId != "" {
		filter["tenantId"], _ = primitive.ObjectIDFromHex(tenantId)
	}
	if name != "" {
		filter["name"] = bson.M{"$regex": name, "$options": "i"}
	}
	if description != "" {
		filter["description"] = bson.M{"$regex": description, "$options": "i"}
	}
	if email != "" {
		filter["email"] = bson.M{"$regex": email, "$options": "i"}
	}
	if contactPersonName != "" {
		filter["contactPersonName"] = bson.M{"$regex": contactPersonName, "$options": "i"}
	}

	total, err := s.vendorCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)
	findOptions.SetSort(bson.D{{Key: "createdDate", Value: -1}})

	cursor, err := s.vendorCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var vendors []model.Vendor
	if err := cursor.All(ctx, &vendors); err != nil {
		return nil, 0, err
	}

	return vendors, total, nil
}

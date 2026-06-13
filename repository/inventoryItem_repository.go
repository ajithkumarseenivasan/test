package repository

import (
	"context"
	"errors"
	"time"
	"user-management/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type InventoryItemRepository interface {
	SaveInventoryItem(item model.InventoryItem) (bool, error)
	GetInventoryItems(tenantId, name, vendorId string, skip int64, limit int64) ([]model.InventoryItemUI, int64, error)
}

type inventoryItemRepository struct {
	inventoryItemCollection *mongo.Collection
}

func NewInventoryItemRepository(client *mongo.Client) InventoryItemRepository {
	collection := client.Database("stratos").Collection("inventoryItem")
	return &inventoryItemRepository{inventoryItemCollection: collection}
}

func (s *inventoryItemRepository) SaveInventoryItem(item model.InventoryItem) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check for duplicate name within the same tenant
	filter := bson.M{
		"name":     item.Name,
		"tenantId": item.TenantID,
	}
	count, err := s.inventoryItemCollection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return false, errors.New("Inventory Item with the same name already exists.")
	}

	_, err = s.inventoryItemCollection.InsertOne(ctx, item)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *inventoryItemRepository) GetInventoryItems(tenantId, name, vendorId string, skip int64, limit int64) ([]model.InventoryItemUI, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	if tenantId != "" {
		filter["tenantId"], _ = primitive.ObjectIDFromHex(tenantId)
	}
	if name != "" {
		filter["name"] = bson.M{"$regex": name, "$options": "i"}
	}
	if vendorId != "" {
		filter["vendors._id"], _ = primitive.ObjectIDFromHex(vendorId)
	}

	total, err := s.inventoryItemCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	pipeline := mongo.Pipeline{
		{{Key: "$match", Value: filter}},

		{{Key: "$lookup", Value: bson.M{
			"from":         "inventoryCategory",
			"localField":   "categoryId",
			"foreignField": "_id",
			"as":           "category",
		}}},

		{{Key: "$unwind", Value: bson.M{
			"path":                       "$category",
			"preserveNullAndEmptyArrays": true,
		}}},

		{{Key: "$lookup", Value: bson.M{
			"from":         "inventoryStorageLocation",
			"localField":   "storageLocationId",
			"foreignField": "_id",
			"as":           "storage",
		}}},

		{{Key: "$unwind", Value: bson.M{
			"path":                       "$storage",
			"preserveNullAndEmptyArrays": true,
		}}},

		{{Key: "$lookup", Value: bson.M{
			"from":         "inventoryInPlantUnit",
			"localField":   "inventoryUnitId",
			"foreignField": "_id",
			"as":           "inplantUnit",
		}}},

		{{Key: "$unwind", Value: bson.M{
			"path":                       "$inplantUnit",
			"preserveNullAndEmptyArrays": true,
		}}},

		{{Key: "$project", Value: bson.M{
			"_id":       1,
			"name":      1,
			"parLevel":  1,
			"tenantId":  1,
			"createdBy": 1,
			"vendors":   1,
			"category":  "$category.name",
			"storage":   "$storage.name",
			"unit":      "$inplantUnit.name",
		}}},

		{{Key: "$sort", Value: bson.M{
			"createdDate": -1,
		}}},

		{{Key: "$skip", Value: skip}},
		{{Key: "$limit", Value: limit}},
	}

	cursor, err := s.inventoryItemCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var inventoryItems []model.InventoryItemUI
	if err := cursor.All(ctx, &inventoryItems); err != nil {
		return nil, 0, err
	}

	return inventoryItems, total, nil
}

package repository

import (
	"context"
	"time"
	"user-management/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type StorageLocationRepository interface {
	SaveLocation(location model.StorageLocation) (bool, error)
	GetLocations(tenantId string, name string, description string, skip int64, limit int64) ([]model.StorageLocation, int64, error)
}

type storageLocationRepository struct {
	locationCollection *mongo.Collection
}

func NewStorageLocationRepository(client *mongo.Client) StorageLocationRepository {
	collection := client.Database("stratos").Collection("inventoryStorageLocation")
	return &storageLocationRepository{locationCollection: collection}
}

func (s *storageLocationRepository) SaveLocation(location model.StorageLocation) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if location.ID.IsZero() {
		location.ID = primitive.NewObjectID()
	}
	if location.CreatedDate.IsZero() {
		location.CreatedDate = time.Now().UTC()
	}
	location.ModifiedDate = time.Now().UTC()

	_, err := s.locationCollection.InsertOne(ctx, location)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *storageLocationRepository) GetLocations(tenantId string, name string, description string, skip int64, limit int64) ([]model.StorageLocation, int64, error) {
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

	total, err := s.locationCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)
	findOptions.SetSort(bson.D{{Key: "createdDate", Value: -1}})

	cursor, err := s.locationCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var locations []model.StorageLocation
	if err := cursor.All(ctx, &locations); err != nil {
		return nil, 0, err
	}

	return locations, total, nil
}

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

type InPlantUnitRepository interface {
	SaveInPlantUnit(inPlantUnit model.InPlantUnit) (bool, error)
	GetInPlantUnits(tenantId string, name string, description string, skip int64, limit int64) ([]model.InPlantUnit, int64, error)
}

type inPlantUnitRepository struct {
	inPlantUnitCollection *mongo.Collection
}

func NewInPlantUnitRepository(client *mongo.Client) InPlantUnitRepository {
	collection := client.Database("stratos").Collection("inventoryInPlantUnit")
	return &inPlantUnitRepository{inPlantUnitCollection: collection}
}

func (i *inPlantUnitRepository) SaveInPlantUnit(inPlantUnit model.InPlantUnit) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Check for duplicate name within the same tenant
	filter := bson.M{
		"name":     inPlantUnit.Name,
		"tenantId": inPlantUnit.TenantID,
	}
	count, err := i.inPlantUnitCollection.CountDocuments(ctx, filter)
	if err != nil {
		return false, err
	}
	if count > 0 {
		return false, errors.New("In-Plant Unit with the same name already exists.")
	}

	if inPlantUnit.ID.IsZero() {
		inPlantUnit.ID = primitive.NewObjectID()
	}

	if inPlantUnit.CreatedDate.IsZero() {
		inPlantUnit.CreatedDate = time.Now().UTC()
	}
	inPlantUnit.ModifiedDate = time.Now().UTC()

	_, err = i.inPlantUnitCollection.InsertOne(ctx, inPlantUnit)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (i *inPlantUnitRepository) GetInPlantUnits(tenantId string, name string, description string, skip int64, limit int64) ([]model.InPlantUnit, int64, error) {
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

	total, err := i.inPlantUnitCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)
	findOptions.SetSort(bson.D{{Key: "createdDate", Value: -1}})

	cursor, err := i.inPlantUnitCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var inPlantUnits []model.InPlantUnit
	if err := cursor.All(ctx, &inPlantUnits); err != nil {
		return nil, 0, err
	}

	return inPlantUnits, total, nil
}

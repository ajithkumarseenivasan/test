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

type CategoryRepository interface {
	SaveCategory(category model.Category) (bool, error)
	GetCategories(tenantId string, skip int64, limit int64) ([]model.Category, int64, error)
}

type categoryRepository struct {
	categoryCollection *mongo.Collection
}

func NewCategoryRepository(client *mongo.Client) CategoryRepository {
	collection := client.Database("stratos").Collection("inventoryCategory")
	return &categoryRepository{categoryCollection: collection}
}

func (c *categoryRepository) SaveCategory(category model.Category) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if category.ID.IsZero() {
		category.ID = primitive.NewObjectID()
	}

	if category.CreatedDate.IsZero() {
		category.CreatedDate = time.Now().UTC()
	}
	category.ModifiedDate = time.Now().UTC()

	_, err := c.categoryCollection.InsertOne(ctx, category)
	if err != nil {
		return false, err
	}

	return true, nil
}

func (c *categoryRepository) GetCategories(tenantId string, skip int64, limit int64) ([]model.Category, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	if tenantId != "" {
		filter["tenantId"] = tenantId
	}

	total, err := c.categoryCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)
	findOptions.SetSort(bson.D{{Key: "createdDate", Value: -1}})

	cursor, err := c.categoryCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var categories []model.Category
	if err := cursor.All(ctx, &categories); err != nil {
		return nil, 0, err
	}

	return categories, total, nil
}

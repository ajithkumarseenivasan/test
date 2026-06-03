package repository

import (
	"context"
	"time"
	"user-management/model"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type CategoryRepository interface {
	SaveCategory(category model.Category) (bool, error)
}

type categoryRepository struct {
	categoryCollection *mongo.Collection
}

func NewCategoryRepository(client *mongo.Client) CategoryRepository {
	collection := client.Database("linga").Collection("category")
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

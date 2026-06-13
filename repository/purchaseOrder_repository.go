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

type PurchaseOrderRepository interface {
	SavePurchaseOrder(item model.PurchaseOrder) (bool, error)
	GetPurchaseOrders(tenantId, status, vendorName, orderId string, skip int64, limit int64) ([]model.PurchaseOrder, int64, error)
}

type purchaseOrderRepository struct {
	purchaseOrderCollection *mongo.Collection
}

func NewPurchaseOrderRepository(client *mongo.Client) PurchaseOrderRepository {
	collection := client.Database("stratos").Collection("purchaseOrders")
	return &purchaseOrderRepository{purchaseOrderCollection: collection}
}

func (s *purchaseOrderRepository) SavePurchaseOrder(item model.PurchaseOrder) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	_, err := s.purchaseOrderCollection.InsertOne(ctx, item)
	if err != nil {
		return false, err
	}
	return true, nil
}

func (s *purchaseOrderRepository) GetPurchaseOrders(tenantId, status, vendorName, orderId string, skip int64, limit int64) ([]model.PurchaseOrder, int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{}
	if tenantId != "" {
		filter["tenantId"], _ = primitive.ObjectIDFromHex(tenantId)
	}
	if vendorName != "" {
		filter["vendorName"] = bson.M{"$regex": vendorName, "$options": "i"}
	}
	if status != "" {
		filter["status"] = bson.M{"$regex": status, "$options": "i"}
	}
	if orderId != "" {
		filter["_id"], _ = primitive.ObjectIDFromHex(orderId)
	}

	total, err := s.purchaseOrderCollection.CountDocuments(ctx, filter)
	if err != nil {
		return nil, 0, err
	}

	findOptions := options.Find()
	findOptions.SetSkip(skip)
	findOptions.SetLimit(limit)
	findOptions.SetSort(bson.D{{Key: "createdDate", Value: -1}})

	cursor, err := s.purchaseOrderCollection.Find(ctx, filter, findOptions)
	if err != nil {
		return nil, 0, err
	}
	defer cursor.Close(ctx)

	var purchaseOrder []model.PurchaseOrder
	if err := cursor.All(ctx, &purchaseOrder); err != nil {
		return nil, 0, err
	}

	return purchaseOrder, total, nil
}

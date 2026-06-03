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

type UserRepository interface {
	GetAll() ([]model.User, error)
	GetUserByName(userName string) (model.User, error)
	SaveNewUser(user model.User) (bool, error)
}

type userRepository struct {
	userCollection *mongo.Collection
}

func NewUserRepository(client *mongo.Client) UserRepository {
	collection := client.Database("linga").Collection("users")
	return &userRepository{
		userCollection: collection,
	}
}

func (u *userRepository) GetAll() ([]model.User, error) {
	var users []model.User

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := u.userCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func (u *userRepository) GetUserByName(userName string) (model.User, error) {
	var resultUser model.User
	if userName == "" {
		return resultUser, errors.New("user name can not be empty")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	filter := bson.M{"name": userName}
	err := u.userCollection.FindOne(ctx, filter).Decode(&resultUser)
	if err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return resultUser, errors.New("user not found")
		}
		return resultUser, err
	}

	return resultUser, nil
}

func (u *userRepository) SaveNewUser(user model.User) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if user.ID.IsZero() {
		user.ID = primitive.NewObjectID()
	}

	_, err := u.userCollection.InsertOne(ctx, user)
	if err != nil {
		return false, err
	}

	return true, nil
}

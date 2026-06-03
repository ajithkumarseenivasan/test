package app

import (
	"log"
	"os"
	"user-management/database"
	"user-management/handler"
	"user-management/repository"
	"user-management/service"

	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	UserHandler     *handler.UserHandler
	CategoryHandler *handler.CategoryHandler
	mongoClient     *mongo.Client
}

func NewApplication() *Application {
	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		mongoUri = "mongodb+srv://ajith:YgnkuVHWKIiKu1t2@cluster0.f8nv8.mongodb.net/linga?retryWrites=true&w=majority"
	}

	client, err := database.NewMongoConnection(mongoUri)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	userRepo := repository.NewUserRepository(client)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	categoryRepo := repository.NewCategoryRepository(client)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	return &Application{
		UserHandler:     userHandler,
		CategoryHandler: categoryHandler,
		mongoClient:     client,
	}
}

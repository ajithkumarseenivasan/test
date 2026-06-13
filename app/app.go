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
	UserHandler          *handler.UserHandler
	CategoryHandler      *handler.CategoryHandler
	AuthHandler          *handler.AuthHandler
	UserService          service.UserService
	StorageHandler       *handler.StorageLocationHandler
	InPlantUnitHandler   *handler.InPlantUnitHandler
	VendorHandler        *handler.VendorHandler
	InventoryItemHandler *handler.InventoryItemHandler
	PurchaseOrderHandler *handler.PurchaseOrderHandler
	mongoClient          *mongo.Client
}

func NewApplication(jwtSecret string) *Application {
	mongoUri := os.Getenv("MONGO_URI")
	if mongoUri == "" {
		mongoUri = "mongodb+srv://ajith:YgnkuVHWKIiKu1t2@cluster0.f8nv8.mongodb.net/stratos?retryWrites=true&w=majority"
	}

	client, err := database.NewMongoConnection(mongoUri)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	userRepo := repository.NewUserRepository(client)
	userService := service.NewUserService(userRepo)
	userHandler := handler.NewUserHandler(userService)

	authService := service.NewAuthService(userRepo, jwtSecret)
	authHandler := handler.NewAuthHandler(authService)

	categoryRepo := repository.NewCategoryRepository(client)
	categoryService := service.NewCategoryService(categoryRepo)
	categoryHandler := handler.NewCategoryHandler(categoryService)

	storageRepo := repository.NewStorageLocationRepository(client)
	storageService := service.NewStorageLocationService(storageRepo)
	storageHandler := handler.NewStorageLocationHandler(storageService)

	inPlantUnitRepo := repository.NewInPlantUnitRepository(client)
	inPlantUnitService := service.NewInPlantUnitService(inPlantUnitRepo)
	inPlantUnitHandler := handler.NewInPlantUnitHandler(inPlantUnitService)

	vendorRepo := repository.NewVendorRepository(client)
	vendorService := service.NewVendorService(vendorRepo)
	vendorHandler := handler.NewVendorHandler(vendorService)

	inventoryItemRepo := repository.NewInventoryItemRepository(client)
	inventoryItemService := service.NewInventoryItemService(inventoryItemRepo)
	inventoryItemHandler := handler.NewInventoryItemHandler(inventoryItemService)

	purchaseOrderRepo := repository.NewPurchaseOrderRepository(client)
	purchaseOrderService := service.NewPurchaseOrderService(purchaseOrderRepo)
	purchaseOrderHandler := handler.NewPurchaseOrderHandler(purchaseOrderService)

	return &Application{
		UserHandler:          userHandler,
		CategoryHandler:      categoryHandler,
		AuthHandler:          authHandler,
		UserService:          userService,
		StorageHandler:       storageHandler,
		InPlantUnitHandler:   inPlantUnitHandler,
		VendorHandler:        vendorHandler,
		InventoryItemHandler: inventoryItemHandler,
		PurchaseOrderHandler: purchaseOrderHandler,
		mongoClient:          client,
	}
}

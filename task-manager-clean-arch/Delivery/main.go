package main

import (
	"context"
	"log"
	"time"

	"task-manager-clean-arch/Delivery/controllers"
	"task-manager-clean-arch/Delivery/routers"
	"task-manager-clean-arch/Infrastructure"
	"task-manager-clean-arch/Repositories"
	"task-manager-clean-arch/Usecases"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func main() {
	// Initialize MongoDB connection
	client, err := initMongoDB()
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}
	defer client.Disconnect(context.Background())

	// Get database and collections
	db := client.Database("taskdb")
	taskCollection := db.Collection("tasks")
	userCollection := db.Collection("users")

	// Initialize infrastructure services
	jwtService := infrastructure.NewJWTService("your_secret_key_change_in_production")
	passwordService := infrastructure.NewPasswordService()

	// Initialize repositories
	taskRepo := repositories.NewTaskRepository(taskCollection)
	userRepo := repositories.NewUserRepository(userCollection)

	// Initialize use cases
	taskUseCase := usecases.NewTaskUseCase(taskRepo)
	userUseCase := usecases.NewUserUseCase(userRepo, passwordService, jwtService)

	// Initialize controllers
	taskController := controllers.NewTaskController(taskUseCase)
	userController := controllers.NewUserController(userUseCase)

	// Setup router
	router := routers.SetupRouter(taskController, userController, jwtService)

	// Start server
	log.Println("Server is running on http://localhost:8080")
	if err := router.Run(":8080"); err != nil {
		log.Fatal("Failed to start server:", err)
	}
}

// initialize the MongoDB connection
func initMongoDB() (*mongo.Client, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		return nil, err
	}

	// Ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		return nil, err
	}

	log.Println("Connected to MongoDB successfully!")
	return client, nil
}

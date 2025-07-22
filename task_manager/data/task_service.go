package data

import (
	"context"
	"errors"
	"log"
	"time"

	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Collection to interact with MongoDB
var taskCollection *mongo.Collection

// Connect to MongoDB once when the program starts
func InitMongoConnection() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		log.Fatal("Failed to create Mongo client:", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal("Failed to connect to MongoDB:", err)
	}

	taskCollection = client.Database("taskdb").Collection("tasks")
	log.Println("✅ Connected to MongoDB and ready!")
}

// Create a new task
func CreateTask(newTask models.Task) models.Task {
	if newTask.Status == "" {
		newTask.Status = models.StatusPending
	}

	_, err := taskCollection.InsertOne(context.TODO(), newTask)
	if err != nil {
		log.Fatal("❌ Failed to insert task:", err)
	}

	return newTask
}

// Get all tasks
func GetTasks() []models.Task {
	cursor, err := taskCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		log.Fatal("❌ Failed to get tasks:", err)
	}
	defer cursor.Close(context.TODO())

	var tasks []models.Task
	if err := cursor.All(context.TODO(), &tasks); err != nil {
		log.Fatal("❌ Failed to decode tasks:", err)
	}

	return tasks
}

// Get one task by ID
func GetTaskByID(id string) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	var task models.Task
	err = taskCollection.FindOne(context.TODO(), bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		return nil, errors.New("task not found")
	}

	return &task, nil
}

// Update a task
func UpdateTask(id string, updated models.Task) (*models.Task, error) {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	update := bson.M{
		"$set": bson.M{
			"title":       updated.Title,
			"description": updated.Description,
			"due_date":    updated.DueDate,
			"status":      updated.Status,
		},
	}

	_, err = taskCollection.UpdateOne(context.TODO(), bson.M{"_id": objID}, update)
	if err != nil {
		return nil, errors.New("failed to update task")
	}

	return GetTaskByID(id)
}

// Delete a task
func DeleteTask(id string) error {
	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	_, err = taskCollection.DeleteOne(context.TODO(), bson.M{"_id": objID})
	if err != nil {
		return errors.New("failed to delete task")
	}

	return nil
}

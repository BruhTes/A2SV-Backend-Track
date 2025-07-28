package repositories

import (
	"context"
	"errors"
	"log"
	"time"

	"task-manager-clean-arch/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type TaskRepositoryImpl struct {
	collection *mongo.Collection
}

func NewTaskRepository(collection *mongo.Collection) domain.TaskRepository {
	return &TaskRepositoryImpl{
		collection: collection,
	}
}

func (r *TaskRepositoryImpl) Create(task *domain.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if task.ID == "" {
		task.ID = primitive.NewObjectID().Hex()
	}

	_, err := r.collection.InsertOne(ctx, task)
	if err != nil {
		log.Printf("Failed to insert task: %v", err)
		return err
	}

	return nil
}

func (r *TaskRepositoryImpl) GetByID(id string) (*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	var task domain.Task
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&task)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("task not found")
		}
		log.Printf("Failed to get task by ID: %v", err)
		return nil, err
	}

	return &task, nil
}

func (r *TaskRepositoryImpl) GetAll() ([]*domain.Task, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		log.Printf("Failed to get tasks: %v", err)
		return nil, err
	}
	defer cursor.Close(ctx)

	var tasks []*domain.Task
	if err := cursor.All(ctx, &tasks); err != nil {
		log.Printf("Failed to decode tasks: %v", err)
		return nil, err
	}

	return tasks, nil
}

func (r *TaskRepositoryImpl) Update(id string, task *domain.Task) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	update := bson.M{
		"$set": bson.M{
			"title":       task.Title,
			"description": task.Description,
			"due_date":    task.DueDate,
			"status":      task.Status,
			"updated_at":  time.Now(),
		},
	}

	result, err := r.collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		log.Printf("Failed to update task: %v", err)
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("task not found")
	}

	return nil
}

func (r *TaskRepositoryImpl) Delete(id string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return errors.New("invalid ID format")
	}

	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		log.Printf("Failed to delete task: %v", err)
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("task not found")
	}

	return nil
} 
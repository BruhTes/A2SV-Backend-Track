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

type UserRepositoryImpl struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) domain.UserRepository {
	return &UserRepositoryImpl{
		collection: collection,
	}
}

func (r *UserRepositoryImpl) Create(user *domain.User) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if user.ID == "" {
		user.ID = primitive.NewObjectID().Hex()
	}

	_, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		log.Printf("Failed to insert user: %v", err)
		return err
	}

	return nil
}

func (r *UserRepositoryImpl) GetByUsername(username string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var user domain.User
	filter := bson.M{"username": username}
	err := r.collection.FindOne(ctx, filter).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		log.Printf("Failed to get user by username: %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) GetByID(id string) (*domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, errors.New("invalid ID format")
	}

	var user domain.User
	err = r.collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, errors.New("user not found")
		}
		log.Printf("Failed to get user by ID: %v", err)
		return nil, err
	}

	return &user, nil
}

func (r *UserRepositoryImpl) ExistsByUsername(username string) (bool, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := r.collection.CountDocuments(ctx, bson.M{"username": username})
	if err != nil {
		log.Printf("Failed to check if username exists: %v", err)
		return false, err
	}

	return count > 0, nil
}

func (r *UserRepositoryImpl) GetUserCount() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	count, err := r.collection.CountDocuments(ctx, bson.M{})
	if err != nil {
		log.Printf("Failed to get user count: %v", err)
		return 0, err
	}

	return count, nil
} 
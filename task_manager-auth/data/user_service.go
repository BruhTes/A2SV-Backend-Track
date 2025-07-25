package data

import (
	"context"
	"errors"
	"task_manager/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection

func InitUserCollection(db *mongo.Database) {
	userCollection = db.Collection("users")
}

func RegisterUser(username, password string) (*models.User, error) {
	count, err := userCollection.CountDocuments(context.TODO(), bson.M{"username": username})
	if err != nil {
		return nil, err
	}
	if count > 0 {
		return nil, errors.New("username already taken")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	userCount, err := userCollection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return nil, err
	}
	role := models.RoleUser
	if userCount == 0 {
		role = models.RoleAdmin
	}

	user := models.User{
		ID:       primitive.NewObjectID().Hex(),
		Username: username,
		Password: string(hashedPassword),
		Role:     role,
	}

	// Insert into database
	_, err = userCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}


func FindUserByUsername(username string) (*models.User, error) {
	var user models.User
	filter := bson.M{"username": username}
	err := userCollection.FindOne(context.TODO(), filter).Decode(&user)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return &user, nil
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


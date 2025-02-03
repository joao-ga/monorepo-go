package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"users/database"
	"users/models"
)

func getUserCollection() *mongo.Collection {
	return database.GetCollection("users")
}

func CreateUser(ctx context.Context, user models.User) (*mongo.InsertOneResult, error) {
	userCollection := getUserCollection()
	return userCollection.InsertOne(ctx, user)
}

func GetUserById(ctx context.Context, id string) (*models.User, error) {
	userCollection := getUserCollection()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	var user models.User
	err = userCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func GetAllUsers(ctx context.Context) ([]models.User, error) {
	userCollection := getUserCollection()
	var users []models.User
	cur, err := userCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var user models.User
		if err := cur.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return users, nil
}

func UpdateUser(ctx context.Context, id string, user bson.M) error {
	userCollection := getUserCollection()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := userCollection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": user})

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

func DeleteUser(ctx context.Context, id string) error {
	userCollection := getUserCollection()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	result, err := userCollection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return err
	}
	if result.DeletedCount == 0 {
		return errors.New("user not found")
	}

	return nil
}

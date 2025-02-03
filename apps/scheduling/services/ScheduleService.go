package services

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"scheduling/database"
	"scheduling/models"
)

func getscheduleCollection() *mongo.Collection {
	return database.GetCollection("schedule")
}

func CreateSchedule(ctx context.Context, schedule models.Schedule) (*mongo.InsertOneResult, error) {
	scheduleCollection := getscheduleCollection()
	return scheduleCollection.InsertOne(ctx, schedule)
}

func GetScheduleByid(ctx context.Context, id string) (*models.Schedule, error) {
	scheduleCollection := getscheduleCollection()
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return nil, err
	}
	var schedule models.Schedule

	err = scheduleCollection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&schedule)
	if err != nil {
		return nil, err
	}

	return &schedule, nil

}

func GetAllSchedules(ctx context.Context) ([]models.Schedule, error) {
	scheduleCollection := getscheduleCollection()
	var scheduleList []models.Schedule
	cur, err := scheduleCollection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	defer cur.Close(ctx)
	for cur.Next(ctx) {
		var schedule models.Schedule
		if err := cur.Decode(&schedule); err != nil {
			return nil, err
		}
		scheduleList = append(scheduleList, schedule)
	}
	if err := cur.Err(); err != nil {
		return nil, err
	}

	return scheduleList, nil
}

func UpdateSchedule(ctx context.Context, id string, schedule bson.M) error {
	scheduleCollection := getscheduleCollection()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := scheduleCollection.UpdateOne(ctx, bson.M{"_id": objectId}, bson.M{"$set": schedule})

	if err != nil {
		return err
	}

	if result.MatchedCount == 0 {
		return errors.New("Schedule not found")
	}

	return nil
}

func DeleteSchedule(ctx context.Context, id string) error {
	scheduleCollection := getscheduleCollection()
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	result, err := scheduleCollection.DeleteOne(ctx, bson.M{"_id": objectId})
	if err != nil {
		return err
	}

	if result.DeletedCount == 0 {
		return errors.New("Schedule not found")
	}

	return nil
}

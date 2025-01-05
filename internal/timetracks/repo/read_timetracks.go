package repo

import (
	"ai-powered-study-planner-backend/internal/timetracks/entity"
	"context"
	"time"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *TimetracksRepo) GetCurrentTimetrack(firebaseUID string) (*entity.Timetrack, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	timetrack := &entity.Timetrack{}
	err := r.Collection.FindOne(ctx, bson.M{"firebase_uid": firebaseUID, "end_time": time.Time{}}).Decode(timetrack)
	if err != nil {
		return nil, err
	}

	return timetrack, nil
}

func (r *TimetracksRepo) GetAllTimetracksOfProfile(firebaseUID string, startTime *time.Time, endTime *time.Time) ([]entity.Timetrack, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"firebase_uid": firebaseUID}
	if startTime != nil {
		filter["start_time"] = bson.M{"$gte": *startTime}
	}
	if endTime != nil {
		filter["end_time"] = bson.M{"$lte": *endTime}
	}

	timetrack := []entity.Timetrack{}
	cursor, err := r.Collection.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	err = cursor.All(ctx, &timetrack)
	if err != nil {
		return nil, err
	}

	return timetrack, nil
}

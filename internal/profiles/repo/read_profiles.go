package repo

import (
	"context"
	"time"

	"ai-powered-study-planner-backend/internal/profiles/entity"

	"go.mongodb.org/mongo-driver/bson"
)

func (r *ProfilesRepo) FindByFirebaseUID(firebaseUID string) (*entity.Profile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	profile := &entity.Profile{}
	err := r.Collection.FindOne(ctx, bson.M{"firebase_uid": firebaseUID}).Decode(profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

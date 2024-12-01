package repo

import (
	"context"
	"time"

	"ai-powered-study-planner-backend/pkg/db"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type ProfilesRepo struct {
	*db.BaseRepo[Profile]
}

func NewProfilesRepo(collection *mongo.Collection) *ProfilesRepo {
	return &ProfilesRepo{
		db.NewBaseRepo[Profile](collection),
	}
}

func (r *ProfilesRepo) FindByFirebaseUID(firebaseUID string) (*Profile, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()

	profile := &Profile{}
	err := r.Collection.FindOne(ctx, bson.M{"firebase_uid": firebaseUID}).Decode(profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

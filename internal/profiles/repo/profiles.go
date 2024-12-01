package repo

import (
	"ai-powered-study-planner-backend/pkg/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProfilesRepo struct {
	db.IBaseRepo[Profile]
}

func NewProfilesRepo(collection *mongo.Collection) *ProfilesRepo {
	return &ProfilesRepo{
		db.NewBaseRepo[Profile](collection),
	}
}

package repo

import (
	"ai-powered-study-planner-backend/internal/profiles/entity"
	"ai-powered-study-planner-backend/pkg/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProfilesRepo struct {
	*db.BaseRepo[entity.Profile]
}

func NewProfilesRepo(collection *mongo.Collection) *ProfilesRepo {
	return &ProfilesRepo{
		db.NewBaseRepo[entity.Profile](collection),
	}
}

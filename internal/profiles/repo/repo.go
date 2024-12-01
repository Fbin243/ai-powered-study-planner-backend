package repo

import (
	"ai-powered-study-planner-backend/internal/profiles/entity"
	"ai-powered-study-planner-backend/pkg/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type profilesRepo struct {
	*db.BaseRepo[entity.Profile]
}

func NewProfilesRepo(collection *mongo.Collection) *profilesRepo {
	return &profilesRepo{
		db.NewBaseRepo[entity.Profile](collection),
	}
}

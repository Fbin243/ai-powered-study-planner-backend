package repo

import (
	"ai-powered-study-planner-backend/internal/timetracks/entity"
	"ai-powered-study-planner-backend/pkg/db"

	"go.mongodb.org/mongo-driver/mongo"
)

type TimetracksRepo struct {
	*db.BaseRepo[entity.Timetrack]
}

func NewTimetracksRepo(col *mongo.Collection) *TimetracksRepo {
	return &TimetracksRepo{db.NewBaseRepo[entity.Timetrack](col)}
}

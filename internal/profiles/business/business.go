package business

import (
	"ai-powered-study-planner-backend/internal/profiles/entity"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type ProfilesRepo interface {
	Insert(profile *entity.Profile) (*entity.Profile, error)
	UpdateById(ID primitive.ObjectID, profile *entity.Profile) (*entity.Profile, error)
	FindByFirebaseUID(firebaseUID string) (*entity.Profile, error)
}

type profilesBusiness struct {
	ProfilesRepo ProfilesRepo
}

func NewProfileBusiness(profilesRepo ProfilesRepo) *profilesBusiness {
	return &profilesBusiness{
		ProfilesRepo: profilesRepo,
	}
}

package business

import (
	"context"

	"ai-powered-study-planner-backend/internal/profiles/repo"
	"ai-powered-study-planner-backend/pkg/auth"
	"ai-powered-study-planner-backend/pkg/db"
	"ai-powered-study-planner-backend/pkg/errors"

	"go.mongodb.org/mongo-driver/mongo"
)

type ProfileBusiness struct {
	ProfilesRepo *repo.ProfilesRepo
}

func NewProfileBusiness(profilesRepo *repo.ProfilesRepo) *ProfileBusiness {
	return &ProfileBusiness{
		ProfilesRepo: profilesRepo,
	}
}

func (b *ProfileBusiness) GetProfile(ctx context.Context) (*repo.Profile, error) {
	firebaseProfile, ok := ctx.Value(auth.ProfileKey).(*auth.FirebaseProfile)
	if !ok {
		return nil, errors.ErrUserUnauthorized
	}

	var profile *repo.Profile
	// Check if the profile already exists
	profile, err := b.ProfilesRepo.FindByFirebaseUID(firebaseProfile.UID)
	if err != nil && err != mongo.ErrNoDocuments {
		return nil, err
	}

	// Create a new profile if it doesn't exist
	if err == mongo.ErrNoDocuments {
		profile = &repo.Profile{
			BaseModel:   &db.BaseModel{},
			FirebaseUID: firebaseProfile.UID,
			Email:       firebaseProfile.Email,
			Name:        firebaseProfile.Name,
			Picture:     firebaseProfile.Picture,
		}

		profile, err = b.ProfilesRepo.Insert(profile)
		if err != nil {
			return nil, err
		}
	}

	return profile, nil
}

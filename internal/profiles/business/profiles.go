package business

import (
	"context"

	"ai-powered-study-planner-backend/internal/profiles/repo"
	"ai-powered-study-planner-backend/pkg/auth"
	"ai-powered-study-planner-backend/pkg/errors"
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

	profile := &repo.Profile{
		Name:        firebaseProfile.Name,
		Email:       firebaseProfile.Email,
		FirebaseUID: firebaseProfile.UID,
		Picture:     firebaseProfile.Picture,
	}

	return profile, nil
}

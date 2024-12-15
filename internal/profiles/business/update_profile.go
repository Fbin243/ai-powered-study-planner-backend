package business

import (
	"context"

	"ai-powered-study-planner-backend/internal/profiles/entity"
	"ai-powered-study-planner-backend/internal/profiles/transport/dto"
	"ai-powered-study-planner-backend/pkg/auth"
	"ai-powered-study-planner-backend/pkg/errors"
)

func (b *ProfilesBusiness) UpdateProfile(ctx context.Context, input *dto.UpdateProfileDto) (*entity.Profile, error) {
	firebaseProfile, ok := ctx.Value(auth.ProfileKey).(*auth.FirebaseProfile)
	if !ok {
		return nil, errors.ErrUserUnauthorized
	}

	// Check the profile exists
	profile, err := b.ProfilesRepo.FindByFirebaseUID(firebaseProfile.UID)
	if err != nil {
		return nil, err
	}

	// Update the profile
	profile.Name = input.Name
	profile.Picture = input.Picture

	profile, err = b.ProfilesRepo.UpdateById(profile.ID, profile)
	if err != nil {
		return nil, err
	}

	return profile, nil
}

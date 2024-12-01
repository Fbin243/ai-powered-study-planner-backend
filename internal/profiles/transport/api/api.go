package api

import (
	"context"

	"ai-powered-study-planner-backend/internal/profiles/entity"
	"ai-powered-study-planner-backend/internal/profiles/transport/dto"
)

type ProfilesBusiness interface {
	GetProfile(ctx context.Context) (*entity.Profile, error)
	UpdateProfile(ctx context.Context, input *dto.UpdateProfileDto) (*entity.Profile, error)
}

type api struct {
	ProfileBusiness ProfilesBusiness
}

func NewAPI(profileBusiness ProfilesBusiness) *api {
	return &api{
		ProfileBusiness: profileBusiness,
	}
}

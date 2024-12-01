package api

import (
	"context"

	"ai-powered-study-planner-backend/internal/profiles/entity"
)

type ProfilesBusiness interface {
	GetProfile(ctx context.Context) (*entity.Profile, error)
}

type api struct {
	ProfileBusiness ProfilesBusiness
}

func NewAPI(profileBusiness ProfilesBusiness) *api {
	return &api{
		ProfileBusiness: profileBusiness,
	}
}

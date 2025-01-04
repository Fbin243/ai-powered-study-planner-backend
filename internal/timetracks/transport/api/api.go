package api

import "ai-powered-study-planner-backend/internal/timetracks/business"

type TimetracksAPI struct {
	TimetracksBusiness *business.TimetracksBusiness
}

func NewTimetracksAPI(timetracksBusiness *business.TimetracksBusiness) *TimetracksAPI {
	return &TimetracksAPI{
		TimetracksBusiness: timetracksBusiness,
	}
}

package api

import "ai-powered-study-planner-backend/internal/analytics/business"

type AnalyticsAPI struct {
	AnalyticsBusiness *business.AnalyticsBusiness
}

func NewAnalyticsAPI(analyticsBusiness *business.AnalyticsBusiness) *AnalyticsAPI {
	return &AnalyticsAPI{
		AnalyticsBusiness: analyticsBusiness,
	}
}

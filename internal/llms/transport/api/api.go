package api

import "ai-powered-study-planner-backend/internal/llms/business"

type LLMsAPI struct {
	LLMsBusiness *business.LLMsBusiness
}

func NewLLMsAPI(llmsBusiness *business.LLMsBusiness) *LLMsAPI {
	return &LLMsAPI{
		LLMsBusiness: llmsBusiness,
	}
}

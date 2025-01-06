package api

import (
	"ai-powered-study-planner-backend/internal/analytics/transport/dto"
	"ai-powered-study-planner-backend/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func (api *LLMsAPI) GiveFeedbacks(c *gin.Context) {
	var filter dto.DateTimeFilter
	if err := c.ShouldBindJSON(&filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	validate := validator.New()
	if err := validate.Struct(filter); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := api.LLMsBusiness.GiveFeedbacks(c.Request.Context(), &filter)
	if err != nil {
		if err == errors.ErrUserUnauthorized {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, result)
}

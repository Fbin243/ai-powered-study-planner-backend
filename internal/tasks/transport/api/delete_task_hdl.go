package api

import (
	"ai-powered-study-planner-backend/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *TasksAPI) DeleteTask(c *gin.Context) {
	taskId := c.Param("id")
	deleteTask, err := api.TasksBusiness.DeleteTask(c.Request.Context(), taskId)
	if err != nil {
		if err == errors.ErrUserUnauthorized {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, deleteTask)
}

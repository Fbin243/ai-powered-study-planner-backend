package api

import (
	"ai-powered-study-planner-backend/pkg/errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

func (api *TasksAPI) GetTasks(c *gin.Context) {
	tasks, err := api.TasksBusiness.GetTasks(c.Request.Context())
	if err != nil {
		if err == errors.ErrUserUnauthorized {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, tasks)
}

func (api *TasksAPI) GetTask(c *gin.Context) {
	taskId := c.Param("id")
	task, err := api.TasksBusiness.GetTaskByID(c.Request.Context(), taskId)
	if err != nil {
		if err == errors.ErrUserUnauthorized {
			c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, task)
}
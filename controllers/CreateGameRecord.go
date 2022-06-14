package controllers

import (
	"gin-glm-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type CreateGameRecordInput struct {
	Name              string `json:"name" binding:"required"`
	Platform          string `json:"platform" binding:"required"`
	Medal             string `json:"medal" binding:"omitempty,oneof=Silver Gold"`
	CompleteTimeHours int    `json:"complete_time_hours" binding:"required,gt=0"`
	Genre             string `json:"genre"`
}

func CreateNewGameRecod(c *gin.Context) {
	var input CreateGameRecordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	record := models.GameRecord{
		Name:              input.Name,
		Platform:          input.Platform,
		Medal:             input.Medal,
		CompleteTimeHours: input.CompleteTimeHours,
		Genre:             input.Genre,
	}
	models.DB.Create(&record)

	c.JSON(http.StatusOK, gin.H{"data": record})
}

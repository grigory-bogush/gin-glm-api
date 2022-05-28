package controllers

import (
	"gin-glm-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UpdateGameRecordInput struct {
	Name              string `json:"name"`
	Platform          string `json:"platform"`
	Medal             string `json:"medal" binding:"omitempty,oneof=Silver Gold"`
	CompleteTimeHours int    `json:"complete_time_hours" binding:"omitempty,gt=0"`
	Genre             string `json:"genre"`
}

func UpdateGameRecord(c *gin.Context) {
	// Get model if exist
	var record models.GameRecord
	if err := models.DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateGameRecordInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&record).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": record})
}

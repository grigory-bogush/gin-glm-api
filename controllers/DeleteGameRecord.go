package controllers

import (
	"gin-glm-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func DeleteGameRecord(c *gin.Context) {
	// Get model if exist
	var record models.GameRecord
	if err := models.DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&record)

	c.JSON(http.StatusOK, gin.H{"data": true})
}

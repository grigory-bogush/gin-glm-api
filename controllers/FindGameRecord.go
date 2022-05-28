package controllers

import (
	"gin-glm-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func FindRecord(c *gin.Context) {
	var record models.GameRecord
	if err := models.DB.Where("id = ?", c.Param("id")).First(&record).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": record})
}

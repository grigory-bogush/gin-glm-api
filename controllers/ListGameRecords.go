package controllers

import (
	"gin-glm-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GET /game-records
// Get all books
func ListGameRecords(c *gin.Context) {
	var game_records []models.GameRecord
	models.DB.Find(&game_records)

	c.JSON(http.StatusOK, gin.H{"data": game_records})
}

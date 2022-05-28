package main

import (
	"gin-glm-api/controllers"
	"gin-glm-api/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": "0.0.1", "status": "Up"})
	})
	r.GET("/game-records", controllers.ListGameRecords)
	r.POST("/create-game-record", controllers.CreateNewGameRecod)
	r.GET("/game-record/:id", controllers.FindRecord)
	r.PATCH("/update-game-record/:id", controllers.UpdateGameRecord)
	r.DELETE("/delete-game-record/:id", controllers.DeleteGameRecord)

	models.ConnectDatabase()

	r.Run()
}

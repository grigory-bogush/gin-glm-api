package main

import (
	"gin-glm-api/controllers"
	"gin-glm-api/models"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	r := gin.Default()

	admin_username := os.Getenv("ADMIN_USERNAME")
	admin_pass := os.Getenv("ADMIN_PASSWORD")

	if admin_username == "" || admin_pass == "" {
		print("No username or password provided")
		os.Exit(1)
	}

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		admin_username: admin_pass,
	}))

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"version": "0.0.2", "status": "Up"})
	})
	authorized.GET("/game-records", controllers.ListGameRecords)
	authorized.POST("/create-game-record", controllers.CreateNewGameRecod)
	authorized.GET("/game-record/:id", controllers.FindRecord)
	authorized.PATCH("/update-game-record/:id", controllers.UpdateGameRecord)
	authorized.DELETE("/delete-game-record/:id", controllers.DeleteGameRecord)

	models.ConnectDatabase()

	r.Run()
}

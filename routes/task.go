package routes

import (
	"github/jashandeep31/todo/database"
	"github/jashandeep31/todo/database/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTaks(c *gin.Context) {
	var tasks []models.Task
	result := database.DB.Find(&tasks)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to fetch tasks",
		})
		return
	}

	c.JSON(200, gin.H{
		"tasks": tasks,
	})
}

func TaskRoutes(r *gin.RouterGroup) {

	r.GET("/tasks", GetAllTaks)
}

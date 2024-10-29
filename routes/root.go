package routes

import "github.com/gin-gonic/gin"

func SetupRoutes(r *gin.Engine) {
	// version of the api
	v1 := r.Group("/api/v1")

	// all tasks route with the v1 group
	TaskRoutes(v1)
	UserRoutes(v1)

	// route to the health of the application
	r.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello world",
		})
	})
}

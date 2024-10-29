package routes

import (
	"github/jashandeep31/todo/database"
	"github/jashandeep31/todo/database/models"
	"github/jashandeep31/todo/validators"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetAllUser(c *gin.Context) {
	var users []models.User
	result := database.DB.Find(&users)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get all users",
		})
		return
	}
	c.JSON(200, gin.H{
		"users": users,
	})
}

func RegisterUser(c *gin.Context) {
	var userInput = &validators.UserInput{}

	if err := c.BindJSON(userInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if validationErrors, err := validators.ValidateInput(userInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": validationErrors,
		})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userInput.Password), 14)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hashpassowrd",
		})
	}

	user := models.User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: string(hashedPassword),
	}
	c.JSON(201, gin.H{
		"user": user,
	})
}

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/users", GetAllUser)
	r.POST("/users/signup", RegisterUser)
}

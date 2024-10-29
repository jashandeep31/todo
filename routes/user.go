package routes

import (
	"fmt"
	"github/jashandeep31/todo/database"
	"github/jashandeep31/todo/database/models"
	"github/jashandeep31/todo/validators"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/jackc/pgx/v5/pgconn"
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
		return
	}

	user := models.User{
		Name:     userInput.Name,
		Email:    userInput.Email,
		Password: string(hashedPassword),
	}

	result := database.DB.Create(&user)
	if result.Error != nil {
		if pgErr, ok := result.Error.(*pgconn.PgError); ok && pgErr.Code == "23505" {
			c.JSON(http.StatusConflict, gin.H{
				"error": "User with this email already exists",
			})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to Generate user",
		})
		return
	}
	c.JSON(201, gin.H{
		"user": user,
	})
}

func LoginUser(c *gin.Context) {
	type input struct {
		Email    string `json:"email" validate:"required,email"`
		Password string `json:"password" validate:"required"`
	}
	var userInput input

	if err := c.BindJSON(&userInput); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}
	if validationErrors, err := validators.ValidateInput(userInput); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"errors": validationErrors,
		})
		return
	}

	var user models.User
	var result = database.DB.First(&user, "email = ? ", userInput.Email)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to get user",
		})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(userInput.Password)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Password is wrong",
		})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"name": user.Name,
		"id":   user.ID,
		// You can add more claims if needed
	})
	signingKey := []byte(os.Getenv("JWT_KEY"))
	s, err := token.SignedString(signingKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to hash the passowrd",
		})
		return
	}

	c.JSON(200, gin.H{
		"user":  user,
		"token": s,
	})
}

func UserRoutes(r *gin.RouterGroup) {
	r.GET("/users", GetAllUser)
	r.POST("/users/signup", RegisterUser)
	r.POST("/users/login", LoginUser)
}

package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/sebas7603/waco-test-go/models"
	"golang.org/x/crypto/bcrypt"
)

var secretKey = []byte(os.Getenv("JWT_SECRET"))

type RegisterInput struct {
	Name      string    `form:"name" binding:"required"`
	Email     string    `form:"email" binding:"required"`
	Password  string    `form:"password" binding:"required"`
	Birthdate time.Time `form:"birthdate" binding:"required" time_format:"2006-01-02"`
	Address   string    `form:"address" binding:"required"`
	City      string    `form:"city" binding:"required"`
}

func Register(c *gin.Context) {
	var input RegisterInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user := models.User{
		Name:      input.Name,
		Email:     input.Email,
		Password:  string(hashedPassword),
		Birthdate: input.Birthdate,
		Address:   input.Address,
		City:      input.City,
	}

	if err := models.CreateUser(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Removing password from response
	user.Password = ""

	token, err := createAuthToken(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "registration success", "user": user, "token": token})
}

func createAuthToken(user *models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(time.Hour * 168).Unix(),
	})

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println("Error generating token:", err)
		return "", err
	}

	return tokenString, nil
}

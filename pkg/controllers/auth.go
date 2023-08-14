package controllers

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sebas7603/waco-test-go/models"
	"golang.org/x/crypto/bcrypt"
)

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

	c.JSON(http.StatusCreated, gin.H{"message": "registration success", "user": user})
}

package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sebas7603/waco-test-go/models"
)

type UpdateInput struct {
	Name      string    `form:"name" binding:"required"`
	Email     string    `form:"email" binding:"required"`
	Birthdate time.Time `form:"birthdate" binding:"required" time_format:"2006-01-02"`
	Address   string    `form:"address" binding:"required"`
	City      string    `form:"city" binding:"required"`
}

func ShowProfile(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	user, err := models.GetUserByID(userID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func UpdateProfile(c *gin.Context) {
	var input UpdateInput

	if err := c.ShouldBind(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	user, err := models.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	user.Name = input.Name
	user.Email = input.Email
	user.Birthdate = input.Birthdate
	user.Address = input.Address
	user.City = input.City

	if err := models.UpdateUser(user); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

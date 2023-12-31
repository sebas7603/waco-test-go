package controllers

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sebas7603/waco-test-go/models"
	"github.com/sebas7603/waco-test-go/pkg/api"
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

	err = addFavoritesToUserStruct(user, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	err = addFavoritesToUserStruct(user, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func AddFavoriteByID(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	user, err := models.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	favorite := models.Favorite{
		UserID: userID,
		RefAPI: c.PostForm("character_id"),
	}

	err = checkFavoriteIsValid(&favorite)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = models.AddFavorite(&favorite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = addFavoritesToUserStruct(user, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"user": user,
	})
}

func RemoveFavoriteByID(c *gin.Context) {
	userID, _ := strconv.ParseInt(c.GetString("user_id"), 10, 64)
	user, err := models.GetUserByID(userID)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	favorite := models.Favorite{
		UserID: userID,
		RefAPI: c.PostForm("character_id"),
	}

	err = models.RemoveFavorite(&favorite)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = addFavoritesToUserStruct(user, c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"user": user,
	})
}

func checkFavoriteIsValid(favorite *models.Favorite) error {
	_, err := api.GetRickAndMortyCharacter(favorite.RefAPI)
	if err != nil {
		return fmt.Errorf("Character with ID %s doesn't exist", favorite.RefAPI)
	}

	existsInDB, _ := models.CheckFavoriteExists(favorite)
	if existsInDB {
		return fmt.Errorf("Character with ID %s is already a Favorite", favorite.RefAPI)

	}
	return nil
}

func addFavoritesToUserStruct(user *models.User, c *gin.Context) error {
	favString, rows, err := models.GetFavoritesStringByUserID(user.ID)
	if err != nil {
		return err
	}

	if rows == 1 {
		var favorites []models.Character
		favorite, _ := api.GetRickAndMortyCharacter(favString)
		favorites = append(favorites, *favorite)
		user.Favorites = &favorites
	} else {
		user.Favorites, _ = api.GetMultipleRickAndMortyCharacters(favString)
	}

	return nil
}

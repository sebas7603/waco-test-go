package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sebas7603/waco-test-go/pkg/api"
)

func IndexCharacters(c *gin.Context) {
	characters, err := api.GetAllRickAndMortyCharacters()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, characters.Results)
}

func ShowCharacter(c *gin.Context) {
	characterID := c.Param("id")
	character, err := api.GetRickAndMortyCharacter(characterID)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, character)
}

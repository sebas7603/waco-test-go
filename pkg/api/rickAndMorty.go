package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sebas7603/waco-test-go/models"
)

var err error

func GetAllRickAndMortyCharacters() (*models.IndexResponse, error) {
	// Making request to API
	requestURL := fmt.Sprintf("%s/character", os.Getenv("RM_API_URL"))
	response, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Decodign reponse in Index struct
	var rickAndMortyIndexResponse models.IndexResponse
	err = json.NewDecoder(response.Body).Decode(&rickAndMortyIndexResponse)
	if err != nil {
		return nil, err
	}

	return &rickAndMortyIndexResponse, nil
}

func GetRickAndMortyCharacter(characterID string) (*models.Character, error) {
	// Making request to API
	requestURL := fmt.Sprintf("%s/character/%s", os.Getenv("RM_API_URL"), characterID)
	response, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Decodign reponse in Character struct
	var rickAndMortyCharacterResponse models.Character
	err = json.NewDecoder(response.Body).Decode(&rickAndMortyCharacterResponse)
	if err != nil {
		return nil, err
	}

	return &rickAndMortyCharacterResponse, nil
}

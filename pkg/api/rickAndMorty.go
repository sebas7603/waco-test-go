package api

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"github.com/sebas7603/waco-test-go/models"
)

var err error

func GetAllRickAndMortyCharacters(page string) (*models.IndexResponse, error) {
	// Making request to API
	requestURL := fmt.Sprintf("%s/character?page=%s", os.Getenv("RM_API_URL"), page)
	response, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Check the status code of the response
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", response.StatusCode)
	}

	// Decodign reponse in Index struct
	var rickAndMortyIndexResponse models.IndexResponse
	err = json.NewDecoder(response.Body).Decode(&rickAndMortyIndexResponse)
	if err != nil {
		return nil, err
	}

	return &rickAndMortyIndexResponse, nil
}

func GetMultipleRickAndMortyCharacters(charactersString string) (*[]models.Character, error) {
	// Making request to API
	requestURL := fmt.Sprintf("%s/character/%s", os.Getenv("RM_API_URL"), charactersString)
	response, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Check the status code of the response
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", response.StatusCode)
	}

	// Decodign reponse in Index struct
	var rickAndMortyMultipleResponse []models.Character
	err = json.NewDecoder(response.Body).Decode(&rickAndMortyMultipleResponse)
	if err != nil {
		return nil, err
	}

	return &rickAndMortyMultipleResponse, nil
}

func GetRickAndMortyCharacter(characterID string) (*models.Character, error) {
	// Making request to API
	requestURL := fmt.Sprintf("%s/character/%s", os.Getenv("RM_API_URL"), characterID)
	response, err := http.Get(requestURL)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	// Check the status code of the response
	if response.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status code: %d", response.StatusCode)
	}

	// Decodign reponse in Character struct
	var rickAndMortyCharacterResponse models.Character
	err = json.NewDecoder(response.Body).Decode(&rickAndMortyCharacterResponse)
	if err != nil {
		return nil, err
	}

	return &rickAndMortyCharacterResponse, nil
}

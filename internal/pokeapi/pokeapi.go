package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

const baseURL string = "https://pokeapi.co/api/v2"

// LocationArea is a location area in the pokemon world.
type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func ListLocationAreas() (LocationAreasResponse, error) {
	url := baseURL + "/location-area"
	res, err := http.Get(url)
	if err != nil {
		return LocationAreasResponse{}, err
	}
	defer res.Body.Close()

	if res.StatusCode > 299 {
		return LocationAreasResponse{}, errors.New("Error: " + res.Status)
	}

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	locationAreasResponse := LocationAreasResponse{}
	err = json.Unmarshal(data, &locationAreasResponse)
	if err != nil {
		return LocationAreasResponse{}, err
	}

	return locationAreasResponse, nil
}

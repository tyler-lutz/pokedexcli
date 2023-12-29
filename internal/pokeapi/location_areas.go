package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
)

type LocationAreasResponse struct {
	Count    int     `json:"count"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"results"`
}

func (c Client) ListLocationAreas(pageURL *string) (LocationAreasResponse, error) {
	url := baseURL + "/location-area"

	if pageURL != nil {
		url = *pageURL
	}

	if val, ok := c.cache.Get(url); ok {
		locationAreasResponse := LocationAreasResponse{}
		err := json.Unmarshal(val, &locationAreasResponse)
		if err != nil {
			return LocationAreasResponse{}, err
		}
		return locationAreasResponse, nil
	}

	res, err := c.httpClient.Get(url)
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

package pokeapi

import (
	"encoding/json"
	"errors"
	"io"
)

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

	c.cache.Add(url, data)

	return locationAreasResponse, nil
}

func (c Client) GetLocation(name string) (Location, error) {
	url := baseURL + "/location-area/" + name

	if val, ok := c.cache.Get(url); ok {
		locationRes := Location{}
		err := json.Unmarshal(val, &locationRes)
		if err != nil {
			return Location{}, err
		}
		return locationRes, nil
	}

	res, err := c.httpClient.Get(url)
	if err != nil {
		return Location{}, err
	}
	defer res.Body.Close()

	data, err := io.ReadAll(res.Body)
	if err != nil {
		return Location{}, err
	}

	locationRes := Location{}
	err = json.Unmarshal(data, &locationRes)
	if err != nil {
		return Location{}, err
	}

	c.cache.Add(url, data)

	return locationRes, nil
}

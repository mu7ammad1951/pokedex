package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ListLocations(path *string) (RespShallowLocations, error) {
	url := baseURL + "/location-area"
	if path != nil {
		url = *path
	}

	data, exists := c.cache.Get(url)
	if !exists {
		fmt.Println("Cache Miss!")
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespShallowLocations{}, err
		}

		LocationsRes, err := c.httpClient.Do(req)
		if err != nil {
			return RespShallowLocations{}, err
		}
		defer LocationsRes.Body.Close()

		data, err = io.ReadAll(LocationsRes.Body)
		if err != nil {
			return RespShallowLocations{}, err
		}

		c.cache.Add(url, data)
	} else {
		fmt.Println("Cache Hit!")
	}

	var listLocations RespShallowLocations

	if err := json.Unmarshal(data, &listLocations); err != nil {
		return RespShallowLocations{}, err
	}

	return listLocations, nil
}

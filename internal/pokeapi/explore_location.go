package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) ExploreLocation(location *string) (RespLocationInfo, error) {
	if location == nil {
		return RespLocationInfo{}, fmt.Errorf("error: location argument missing 'explore <location-area>'")
	}

	url := baseURL + "/location-area/" + *location

	data, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return RespLocationInfo{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return RespLocationInfo{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return RespLocationInfo{}, err
		}

		c.cache.Add(url, data)
	}

	var locationInfo RespLocationInfo
	if err := json.Unmarshal(data, &locationInfo); err != nil {
		return RespLocationInfo{}, err
	}

	return locationInfo, nil
}

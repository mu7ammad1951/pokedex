package pokeapi

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func (c *Client) GetPokemonInfo(name *string) (Pokemon, error) {

	if name == nil {
		return Pokemon{}, fmt.Errorf("error: pokemon argument missing 'catch <pokemon-name>'")
	}
	url := baseURL + "/pokemon/" + *name
	data, exists := c.cache.Get(url)
	if !exists {
		req, err := http.NewRequest("GET", url, nil)
		if err != nil {
			return Pokemon{}, err
		}

		res, err := c.httpClient.Do(req)
		if err != nil {
			return Pokemon{}, err
		}
		defer res.Body.Close()

		data, err = io.ReadAll(res.Body)
		if err != nil {
			return Pokemon{}, err
		}
	}
	var pokemonInfo Pokemon
	if err := json.Unmarshal(data, &pokemonInfo); err != nil {
		return Pokemon{}, err
	}
	return pokemonInfo, nil
}

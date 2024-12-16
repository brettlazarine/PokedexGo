package main

import (
	"encoding/json"
	"fmt"
	"pokedexgo/internal/api"
)

func commandMap(configPtr *Config, arg ...string) error {
	if configPtr.Next == "" {
		configPtr.Next = "https://pokeapi.co/api/v2/location-area"
	}

	if cached, ok := cache.Get(configPtr.Next); ok {
		var areas api.PokeArea
		err := json.Unmarshal(cached, &areas)
		if err != nil {
			return err
		}

		configPtr.Next = areas.Next
		configPtr.Previous = areas.Previous

		for _, area := range areas.Results {
			fmt.Println(area.Name)
		}

		return nil
	}

	apiClient := api.Api{}
	res, err := apiClient.Get(configPtr.Next)
	if err != nil {
		return err
	}
	cache.Add(configPtr.Next, res)

	var areas api.PokeArea
	err = json.Unmarshal(res, &areas)
	if err != nil {
		return err
	}

	configPtr.Next = areas.Next
	configPtr.Previous = areas.Previous

	for _, area := range areas.Results {
		fmt.Println(area.Name)
	}

	return nil
}

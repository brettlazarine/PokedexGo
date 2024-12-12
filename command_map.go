package main

import (
	"encoding/json"
	"fmt"
	"pokedexgo/internal/api"
)

func commandMap(configPtr *Config) error {
	if configPtr.Next == "" {
		configPtr.Next = "https://pokeapi.co/api/v2/location-area"
	}
	apiClient := api.Api{}
	res, err := apiClient.Get(configPtr.Next)
	if err != nil {
		return err
	}

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

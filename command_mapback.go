package main

import (
	"encoding/json"
	"fmt"
	"pokedexgo/internal/api"
)

func commandMapback(configPtr *Config) error {
	if configPtr.Previous == nil || *configPtr.Previous == "" {
		fmt.Println("Currently on the first page")
		return nil
	}

	apiClient := api.Api{}
	res, err := apiClient.Get(*configPtr.Previous)
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

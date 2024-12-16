package main

import (
	"encoding/json"
	"fmt"
	"pokedexgo/internal/api"
)

func commandMapback(configPtr *Config, arg ...string) error {
	if configPtr.Previous == nil || *configPtr.Previous == "" {
		fmt.Println("Currently on the first page")
		return nil
	}

	if cached, ok := cache.Get(*configPtr.Previous); ok {
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
	res, err := apiClient.Get(*configPtr.Previous)
	if err != nil {
		return err
	}
	cache.Add(*configPtr.Previous, res)

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

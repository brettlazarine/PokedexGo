package main

import (
	"encoding/json"
	"fmt"
	"pokedexgo/internal/api"
)

func commandExplore(configPtr *Config, name ...string) error {
	url := "https://pokeapi.co/api/v2/location-area/" + name[0]
	if name == nil || len(name) == 0 {
		return fmt.Errorf("please provide an area name or id")
	}
	var explorePokemon api.ExplorePokemon

	if cached, ok := cache.Get(url); ok {
		err := json.Unmarshal(cached, &explorePokemon)
		if err != nil {
			return err
		}

	} else {
		apiClient := api.Api{}
		res, err := apiClient.Get(url)
		if err != nil {
			return err
		}
		cache.Add(url, res)

		err = json.Unmarshal(res, &explorePokemon)
		if err != nil {
			return err
		}
	}

	for _, pokemon := range explorePokemon.PokemonEncounters {
		fmt.Println(pokemon.Pokemon.Name)
	}

	return nil
}

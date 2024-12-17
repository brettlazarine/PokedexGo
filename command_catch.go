package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"pokedexgo/internal/api"
)

func commandCatch(configPtr *Config, name ...string) error {
	if len(name) == 0 {
		return fmt.Errorf("please provide a pokemon name")
	}
	url := "https://pokeapi.co/api/v2/pokemon/" + name[0]

	if _, ok := pokedex[name[0]]; ok {
		return fmt.Errorf("pokemon already in pokedex")
	}

	apiClient := api.Api{}
	res, err := apiClient.Get(url)
	if err != nil {
		return err
	}

	var pokemon api.Pokemon
	err = json.Unmarshal(res, &pokemon)
	if err != nil {
		return err
	}

	fmt.Printf("Throwing a Pokeball at %v...", name[0])

	difficulty := rand.Intn(pokemon.BaseExperience)
	if difficulty > 100 {
		return fmt.Errorf("pokemon got away")
	}

	fmt.Printf("%v was caught!\n", name[0])

	pokedex[name[0]] = pokemon

	return nil
}

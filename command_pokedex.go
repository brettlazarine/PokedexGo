package main

import (
	"fmt"
)

func commandPokedex(configPtr *Config, arg ...string) error {
	if pokedex == nil {
		return fmt.Errorf("no pokemon in the pokedex")
	}

	for _, pokemon := range pokedex {
		fmt.Printf("- %s\n", pokemon.Name)
	}

	return nil
}

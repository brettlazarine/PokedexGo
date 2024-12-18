package main

import (
	"fmt"
)

func commandInspect(configPtr *Config, name ...string) error {
	if len(name) == 0 {
		return fmt.Errorf("please provide a pokemon name")
	}
	if pokemon, ok := pokedex[name[0]]; ok {
		fmt.Printf("Name: %v\n", pokemon.Name)
		fmt.Printf("Height: %v\n", pokemon.Height)
		fmt.Printf("Weight: %v\n", pokemon.Weight)
		fmt.Printf("Stats:\n")
		for _, value := range pokemon.Stats {
			fmt.Printf("  -%v: %v\n", value.Stat.Name, value.BaseStat)
		}
		fmt.Printf("Types:\n")
		for _, t := range pokemon.Types {
			fmt.Printf("  - %v\n", t.Type.Name)
		}

	} else {
		fmt.Printf("you have not caught that pokemon\n")
		return nil
	}

	return nil
}

// Template output
// Pokedex > inspect pidgey
// you have not caught that pokemon
// Pokedex > catch pidgey
// Throwing a Pokeball at pidgey...
// pidgey was caught!
// Pokedex > inspect pidgey
// Name: pidgey
// Height: 3
// Weight: 18
// Stats:
//   -hp: 40
//   -attack: 45
//   -defense: 40
//   -special-attack: 35
//   -special-defense: 35
//   -speed: 56
// Types:
//   - normal
//   - flying

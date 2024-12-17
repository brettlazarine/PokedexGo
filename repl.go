package main

import (
	"bufio"
	"fmt"
	"os"
	"pokedexgo/internal/api"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(configPtr *Config, arg ...string) error
}
type Config struct {
	Next     string
	Previous *string
}

var commands map[string]cliCommand
var configPtr = &Config{}
var pokedex = make(map[string]api.Pokemon)

func init() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display available commands",
			callback:    commandHelp,
		},
		"map": {
			name:        "map",
			description: "Displays the names of 20 areas in the Pokemon world. ",
			callback:    commandMap,
		},
		"mapb": {
			name:        "mapb",
			description: "Displays the names of the previous 20 areas in the Pokemon world. ",
			callback:    commandMapback,
		},
		"exit": {
			name:        "exit",
			description: "Exit the Pokedex",
			callback:    commandExit,
		},
		"explore": {
			name:        "explore",
			description: "Displays the names of Pokemon that can be found in a specific area. --explore <area> ",
			callback:    commandExplore,
		},
		"catch": {
			name:        "catch",
			description: "Catch a Pokemon. --catch <pokemon name>",
			callback:    commandCatch,
		},
	}
}

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		input := ""

		scanner.Scan()

		words := cleanInputString(scanner.Text())
		if len(words) == 0 {
			continue
		}
		if len(words) > 1 {
			input = words[1]
		}

		command := words[0]
		if cmd, ok := commands[command]; ok {
			err := cmd.callback(configPtr, input)
			if err != nil {
				fmt.Println(err)
			}
			continue
		} else {
			fmt.Println("Command not found. Please try again.")
			continue
		}
	}
}

// Splits user string into slice of words, lowercases input and trims leading and trailing spaces
func cleanInputString(text string) []string {
	if text == "" {
		return []string{}
	}

	words := strings.Fields(text)
	for i := range words {
		words[i] = strings.ToLower(words[i])
		words[i] = strings.Trim((words[i]), " ")
	}
	return words
}

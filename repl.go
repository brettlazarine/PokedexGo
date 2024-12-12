package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type cliCommand struct {
	name        string
	description string
	callback    func(configPtr *Config) error
}
type Config struct {
	Next     string
	Previous *string
}

var commands map[string]cliCommand
var configPtr = &Config{}

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
	}
}

func startREPL() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		fmt.Print("Pokedex > ")
		scanner.Scan()

		words := cleanInputString(scanner.Text())
		if len(words) == 0 {
			continue
		}
		command := words[0]
		if cmd, ok := commands[command]; ok {
			err := cmd.callback(configPtr)
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

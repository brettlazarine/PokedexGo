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
	callback    func() error
}

var commands map[string]cliCommand

func init() {
	commands = map[string]cliCommand{
		"help": {
			name:        "help",
			description: "Display available commands",
			callback:    commandHelp,
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
			err := cmd.callback()
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

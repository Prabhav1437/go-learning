package main

import (
	"fmt"
	"os"
)

type clicommands struct{
	name string
	description string
	callback func()
}

func getcommands() map[string]clicommands{
	return map[string]clicommands{
		"help": {
			name: "help",
			description: "Shows all the available commands",
			callback: func(){
				fmt.Println("Available commands:")
				availablecommands := getcommands()
				for _, command := range availablecommands {
					fmt.Println(">",command.name, "-", command.description)
				}
			},
		},
		"exit": {
			name: "exit",
			description: "Kills the application",
			callback: func(){
				os.Exit(0)
			},
		},
	}
}
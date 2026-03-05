package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func dorepl(){
	for{
		fmt.Print(">>> ")
		meow := bufio.NewScanner(os.Stdin)
		meow.Scan()
		v := meow.Text()
		k := standard(v)
		
		availablecommands := getcommands()
		
		command , ok := availablecommands[k[0]]

		if !ok{
			fmt.Println("Invalid Command")
			continue
		}

		command.callback()
	}
}

func standard(text string) []string{
	l := strings.ToLower(text)
	return strings.Fields(l)
}
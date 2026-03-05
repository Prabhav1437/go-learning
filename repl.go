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
		
		switch k[0]{
		case "exit":
			os.Exit(0)
		case "help":
			fmt.Println("Available commands: exit, help")
		default:
			fmt.Println("Unknown command")
		}
	}
}

func standard(text string) []string{
	l := strings.ToLower(text)
	return strings.Fields(l)
}
package main

import (
	"bufio"
	"fmt"
	"os"
)

func dorepl(){
	for{
		fmt.Print(">>> ")
		meow := bufio.NewScanner(os.Stdin)
		meow.Scan()
		v := meow.Text()
		fmt.Println(v)
	}
}
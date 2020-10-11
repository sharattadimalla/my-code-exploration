package main

import (
	"fmt"
	"os"
	)

func main() {

	argsWithoutProg := os.Args[1:]
	fmt.Println("Command Line args: ", argsWithoutProg)
}

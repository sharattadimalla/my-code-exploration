package main

import (
	"fmt"
	)

func main() {

	if 4%2 == 0 {
		fmt.Println("4 is an even number")
	}

	if 3%2 == 0 {
		fmt.Println("3 is an even number")
	} else {
		fmt.Println("3 is an odd number")
	}

	if -1 < 0 {
		fmt.Println("Number is a negative integer")
	} else if 1 == 1 {
		fmt.Println("Number is equals integer")
	} else {
		fmt.Println("else")
	}
}

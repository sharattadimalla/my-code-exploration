package main

import (
	"fmt"
	)

func main() {

	a := 5
	for a >= 0 {
		fmt.Println(a)
		a = a - 1
	}

	for b := 6; b <= 10; b++ {

		fmt.Println(b)
	}

	for {
		fmt.Println("Inside for loop")
		break
	}	
}	

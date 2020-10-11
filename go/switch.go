package main

import (
	"fmt"
	"time"
	)

func main() {

	i := 1
	fmt.Println("value i ",i)
	switch i {
	case 1:
		fmt.Println("case 1")
	case 2:
		fmt.Println("case 2")
	default:
		fmt.Println("default")
	}
	
	fmt.Println(time.Now().Weekday())
	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("Hurray! It's the Weekend!")
	default:
		fmt.Println("Time to chug along in the week")
	}

	fmt.Println(time.Now())
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon!")
	default:
		fmt.Println("It's after noon!")
	} 	
}

package main

import (
	"fmt"
	)


func main() {

	nums := []int{2,4,6}
	sum := 0
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum : ",sum)

	// range index
	for idx, num := range nums {
		if num == 4 {
			fmt.Println("Position Index for 4 is: ", idx)
		}
	}

	// range over map
	kvs := map[string]string{"a":"avengers", "b":"batman"}
	for k,v := range kvs {
		fmt.Printf("%s -> %s\n", k,v)
	}

	// range over keys
	for k := range kvs {
		fmt.Println("key: ", k)
	}	

	// range over string
	for i,c := range "go" {
		fmt.Println(i,c)
	}
}


package main

import (
	"fmt"
	)

func main() {

	sample_maps := make(map[string]string)
	sample_maps["a"] = "a"
	sample_maps["b"] = "b"

	fmt.Println("map", sample_maps)

	fmt.Println("length of map", len(sample_maps))

	delete (sample_maps, "b")
	fmt.Println("map", sample_maps)

	n := map[string]string{"foo":"foo_val", "bar":"bar_val"}
	fmt.Println(n)
}


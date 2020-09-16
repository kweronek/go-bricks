package main

import (
	"fmt"
	"time"
)

func main() {
	// demo for ranges
	fmt.Println("Demo for ranges")

	a := []string{0: "a", 2: "c", 1: "b"}

	// loop over an array/a slice
	fmt.Println("\nindex, element")
	for i, e := range a {
		// i is the index, e the element
		fmt.Println(i, e)
	}
	// if you only need e:
	fmt.Println("element")
	for _, e := range a {
		// e is the element
		fmt.Println(e)
	}

	// ...and if you only need the index
	fmt.Println("index:")
	for i := range a {
		fmt.Println(i)
	}

	for range time.Tick(time.Second) {
		// do it once a sec
		fmt.Println(time.Now())
	}

}

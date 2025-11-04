package main

import (
"fmt"
)

func main() {
	// Demo for Arrays
	fmt.Println("Demo for Arrays")

	var a [5]int64                  // Array mit 5 Feldern
	fmt.Println("a =", a)
	var b = [3]int{1, 2, 3}         // entspricht [1, 2, 3]
	fmt.Println("b =", b)
	var c = [...]int{1, 7, 4, 8, 3} // automatische Längenerkennung
	fmt.Println("c =", c)
	fmt.Println("c[0:3]=", c[0:3])
	fmt.Println("c[3:5]=", c[3:5])
	c[2]=22
	fmt.Println("c =", c)
}


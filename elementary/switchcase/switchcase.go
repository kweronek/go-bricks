package main

import (
	"fmt"
)

func main() {
	// Demo for switch case
	fmt.Println("Demo for switch-case:")

	var m int = 11

	switch m {
	default: fmt.Println(m, "is a number larger than 9")
	case 0, 2, 4, 6, 8: fmt.Println("lower than 10 and even")
	case 1, 3, 5, 7, 9: fmt.Println("lower than 19 and odd")
	}

	// Switch ohne Variable
	fmt.Println("\nDemo for switch-case without variable:")
	var a int = 7
	var b int = 7
	var c int = 6
	var d int = 7

	switch {
	default: fmt.Println("default")
	case a < b: fmt.Println("( a =", a,") < ( b =",b,")")
	case a > c: fmt.Println("( a =", a,") > ( c =",c,")")
	case b == d: fmt.Println("( b =",b,") == ( d =",d,")")
	}

}


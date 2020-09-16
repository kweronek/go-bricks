package main

import (
	"fmt"
)

func main() {
	// Demo for Loops
	fmt.Println("Demo for Loops:")

	// for Schleife
	fmt.Println("\nfor-Schleife:")
	fmt.Println("for i := 1; i <= 3; i++ {...}")
	summe := 0
	for i := 1; i <= 3; i++ {
		summe += i
		fmt.Println("i, summe:", i, summe)
	}

	fmt.Println("\nwhile-Schleife")
	fmt.Println("for i<4 {...}")
	var i = 0 // while
	for i < 4 {
		fmt.Println("i =", i)
		i = i + 1
	}

	fmt.Println("\ndo-while-Schleife")
	i = 0 //do while (go style)
	for ok := true; ok; ok = i < 3 {
		fmt.Println(i * 10)
		i = i + 1
	}

	i = 0 //repeat until (traditional style)
	for {
		fmt.Println(i * 1000)
		i = i + 1
		if i == 3 {
			break
		}
	}

	i = 0 //repeat until (mixed style)
	for until := false; !until; {
		fmt.Println(i * 1000)
		i = i + 1
		until = (i == 3)
	}

	i = 0 //repeat until (go style)
	for until := false; !until; until = i == 3 {
		fmt.Println(i * 1000)
		i = i + 1
	}

	i = 0
	//do while
	for {
		i = i + 1
		fmt.Println(i * 7)
		if !(i <= 3) {
			break
		}
	}

	i = 0
	//do while
	for while := true; while; {
		i = i + 1
		fmt.Println(i * 2)
		while = i <= 3
	}

	i = 0
	//do while
	//  for ok := true; ok; ok = (i<=3) {
	for ok := true; ok; ok = i <= 3 {
		i = i + 1
		fmt.Println(i * 10)
	}
}

package main

import "fmt"

func main() {
	// Demo for if  and if else
	//if
	var a int = -4
	fmt.Println("a =",a)
	if a < 0 { a = -a }
	fmt.Println(a, "= |a|")
	fmt.Println()

	// if-else
	a=4
	var b int = 4
	fmt.Println("a =", a ,"und b =",b)
	if a==b {
		fmt.Println("a = b")
	} else if a > b {
		fmt.Println("a > b")
	} else {
		fmt.Println("a < b")
	}
}

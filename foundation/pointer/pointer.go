package main

import "fmt"

func main() {
	var v int  // v ist int
	var z *int // z ist Zeiger auf int

	v = 15 // setze a auf 15
	fmt.Println("let v =", v)

	z = &v                     // b ist nun Zeiger auf a
	fmt.Println("z = &v =", z) // zeigt eine Hex-Zahl als Speicheradresse von a
	fmt.Println("*z =", *z)    // *b referenziert auf den Inhalt der Speicheradresse

	fmt.Println("let *z = 20")
	*z = 20                     // setzt den Inhalt der Speicheradresse auf 20
	fmt.Println("*z = v =", *z) // damit ist nun v auch 20

}

package main

import (
	"fmt"
)

// Demo for Methods

// Structs see Demo for Structs
type Person struct {
	ID int
	Vorname string
	Nachname string
	Alter int
}

// Methoden
func (p Person) FullName() string {
		return p.Vorname + " " + p.Nachname
}

func main() {
	fmt.Println("Methoden mit Struct")

	// Deklarationen
	a := Person{1, "Hans", "Mustermann", 56,}
	b := Person{ // ID bleibt 0
			Vorname:  "Erika",
			Nachname: "Musterfrau",
			Alter:    35,
	}
	c := Person{ 2, "Kim", "Muster" , 45 }
	fmt.Println("\n",a, "\n", b, "\n", c,"\n")
	// ergibt: {1 Hans Mustermann 56} {0 Erika Mustermann 35}

	fmt.Println("a.FullName() -->")
	fmt.Println(a.FullName())
	// ergibt: Hans Mustermann
}



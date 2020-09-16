package main

import "fmt"

func main() {
	// structs
	fmt.Println("Verbundtypen mit struct")

	type Person struct {
		ID       int
		Vorname  string
		Nachname string
		Alter    int
	}

	// Deklarationen
	a := Person{1, "Hans", "Mustermann", 56}

	b := Person{ // ID bleibt 0
		Vorname:  "Erika",
		Nachname: "Mustermann",
		Alter:    35,
	}
	c := Person{2, "Kim", "Muster", 45}
	fmt.Println(a, b, c)

	// ergibt:
	// {1 Harald Weidner 56} {0 Erika Mustermann 35}

}

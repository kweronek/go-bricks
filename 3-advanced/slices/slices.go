package main

import (
	"fmt"
)

func main() {
	// Demo for slices
	fmt.Println("Demo: slices")

	var s []int // uninitialisierter Slice
	fmt.Println("s =", s)
	var t = make([]int32, 5) // Slice der Länge 5
	fmt.Println("t =", t)
	var r = make([]int32, 5, 10) // Slice der Länge 5 und Kapazität 10
	fmt.Println("r =", r)

	var c = [...]int{1, 7, 4, 8, 3} // automatische Längenerkennung
	fmt.Println("c =", c)
	s = c[1:4] // s ist [7, 4, 8]
	fmt.Println("c[1:4] =", s)

	r = append(r, 9) // Anhängen eines Elementes, len(r) ist jetzt 6
	fmt.Println("append(r, 9) =", r)
	fmt.Println("len(r) =", len(r))
	t = append(t, 1) // Kapazitätserweiterung nötig, t wird umkopiert
	fmt.Println("append(t, 1) =", t)
	fmt.Println("len(t) =", len(t))
}

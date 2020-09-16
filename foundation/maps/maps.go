package main

import "fmt"

func main() {
	// maps
	fmt.Println("Demo for maps:")
	var m map[string]int
	m = make(map[string]int)
	//	m := make(map[string]int,10) 				// Optional mit Kapazität

	m["Question"] = 99 // Value zuweisen
	m["Answer"] = 42   // Value zuweisen
	fmt.Println("The value of Question:", m["Question"])
	fmt.Println("The value of Answer  :", m["Answer"])

	delete(m, "Answer") //Element einer Map löschen
	fmt.Println("The value of Answer after Deletion:", m["Answer"])

	v, ok := m["Answer"] // Enthält den Value,
	// ok ist true oder false, falls Element nicht vorhanden
	fmt.Println("The value of Answer after Deletion:", v, "is present?", ok)

	fmt.Println("\nAufruf der Map mittels Zeiger:")
	pm := &m                       // Erzeuge einen Pointer auf die Map
	fmt.Println((*pm)["Question"]) // Achtung! Klammern notwendig …
	// … *pm["Answer"] wäre unzulässig bzw.
	// wäre gleichbedeutend mit *(pm["Answer"])

	fmt.Println("\nVerschachtelte Maps:")
	mdm := map[string]map[string]map[int]int{ // Multidimensionale Maps sind möglich
		"color": {"green": {1: 0x00ff00,
			2: 0x00ee00,
		}}}
	fmt.Println("The map mdm: ", mdm)

	// var mdm2 map[string]map[string]map[int]int
	// mdm2 = make(map[string]map[string]map[int]int)
	// oder
	mdm2["color"] = make(map[string]map[int]int)
	mdm2["color"]["red"] = make(map[int]int)
	mdm2["color"]["red"][1] = 0xff0000
	fmt.Println("The map mdm2 :", mdm2["color"]["red"][1])
	fmt.Println(0xff0000)

	mi := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}
	fmt.Println("\nThe map mi: ", mi)
}

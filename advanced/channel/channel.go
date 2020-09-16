package main

import "fmt"
	// Demo for channels

	// chan   // bidirektional
	// chan<- // nur senden
	// <-chan // nur empfangen

	func fibonacci(funcChannel chan int) {
		var f0, f1 int = 0, 1			// Startwerte
		for {
			funcChannel <- f0			// sende f0 an funcChannel
			f0, f1 = f1, f0+f1			// Berechnung
		}
	}

	func main() {
		fmt.Println("Demo for channels")

		myChannel := make(chan int)		// öffnet myChannel (bidir)
		defer close(myChannel)			// schließt den myChannel nach Programmende

		go fibonacci(myChannel)			// startet die fibonacci als thread

		var res int
		for i:=0; i<25; i++ {
			res = <- myChannel			// empfange von myChannel und weise res zu
			fmt.Println(res)
		}
	}

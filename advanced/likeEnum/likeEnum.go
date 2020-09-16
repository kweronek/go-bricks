package main

import (
	"fmt"
)

type direction int

const (
	Nord direction = iota
	Ost
	Süd
	West
)

func (fd direction) String() string {
	return [...]string{"North", "East", "South", "West"}[fd]
}

func main() {
	// like an enumeration
	var fd direction = Nord
	fmt.Println(fd)
	switch fd {
	case Nord:
		fmt.Println("nach oben")
	case Ost:
		fmt.Println("nach rechts")
	case West:
		fmt.Println("nach links")
	case Süd:
		fmt.Println("nach unten")
	default:
		fmt.Println("nach ?")
	}
}


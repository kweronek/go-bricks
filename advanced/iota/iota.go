package main
// demo for iota
import (
	"fmt"
)

func main() {
	fmt.Println("Demo for iota")
	const (
		C0 = iota + 1
		C1
		C2
	)
	fmt.Println(C0,C1,C2)
	const (
		D0 = iota
		D1
		_
		D2
	)
	fmt.Println(D0,D1,D2)
	fmt.Println(C0,C1,C2)
}

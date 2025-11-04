package main

// demo for func and error
import "fmt"

func mult(a, b int64) int64 {
	return a * b
}

func add(a, b int64) (c int64) {
	c = a + b
	return
}

func foo(x float64) (float64, error) {
	y := 1/x
	return y, nil
}

func main() {
	x := mult(3, 8)
	fmt.Println("x := mult(3, 8) =",x)

	y := add(x, 10)
	fmt.Println("y := add(x,10)  =",y)

	var z float64
	var err error
	z, err = foo(7.0)
	if err != nil {
		panic(err)} else {
		fmt.Println(z)
	}
}
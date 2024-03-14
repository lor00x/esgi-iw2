package main

import "fmt"

// Type func(int, int) int
func Add(a, b int) int {
	return a + b
}

var Sub = Operation(func(x, z int) int {
	return x - z
})

// Type nommé
type Operation func(int, int) int

// Operation et func(int, int) int
// sont compatibles
var defaultOp Operation = Add

func main() {
	// Appel de la fonction assignée
	// à la variable defaultOp
	defaultOp = Sub
	fmt.Printf("%#v")
	res := defaultOp(42, 90)
	fmt.Println(res)
}

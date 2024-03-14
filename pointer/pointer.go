package main

import "fmt"

func main() {
	// var age float64 = 42.48

	var pAge *int // = &age

	// Les deux opérations ci-dessous sont équivalentes
	// age = 99
	// *pAge = 99

	fmt.Printf("%d", pAge)

}

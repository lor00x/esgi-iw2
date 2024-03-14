package main

import "fmt"

func main() {

	hi := 4

	// Reuse "hi", create variable "ho"
	hi, ho := 5, 6

	fmt.Println(hi, ho)
}

func Test() (int, error) {
	return 0, nil
}

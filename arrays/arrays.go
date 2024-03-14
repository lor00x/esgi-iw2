package main

import (
	"fmt"
	"slices"
)

const height = 2160
const width = 3840

var image4k [height][width]int

func main() {
	var numbers = []int{}

	fmt.Printf("Len: %d, Cap: %d\n", len(numbers), cap(numbers))

	numbers = append(numbers, 4, 6, 7)
	fmt.Printf("Len: %d, Cap: %d\n", len(numbers), cap(numbers))

	numbers = append(numbers, numbers...)
	fmt.Printf("Len: %d, Cap: %d\n", len(numbers), cap(numbers))
	//clear(numbers)

	slices.Sort(numbers)
	fmt.Println(numbers)
}

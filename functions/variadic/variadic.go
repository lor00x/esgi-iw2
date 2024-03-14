package main

import "fmt"

func main() {
	Sum(1, 3, 5, 7, 9)
}

func Sum(v ...int) int {
	fmt.Printf("%#v", v)
	for index, value := range v {
		fmt.Println(index, value)
	}
	return 0
}

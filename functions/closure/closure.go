package main

import "fmt"

type Generator func() int

func NewGenerator() Generator {
	i := 0 // Etat interne
	return func() int {
		i++
		return i
	}
}

func main() {
	next := NewGenerator()
	one := next()
	two := next()
	three := next()
	fmt.Println(one, two, three, next(), next())
}

package main

import (
	"fmt"
)

type BoxInt struct {
	content int
}

type BoxString struct {
	content string
}

type Box[T any] struct {
	content T
}

func (b Box[T]) Content() T {
	return b.content
}

type User struct {
	Name string
}

func main() {
	b := Box[string]{content: "Hello"}
	b2 := Box[int]{content: 12}
	b3 := Box[User]{content: User{Name: "John"}}
	fmt.Println("Content:", b.content)
	fmt.Println("Content:", b2.content)
	fmt.Println("Content:", b3.content)
}

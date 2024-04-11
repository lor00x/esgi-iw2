package main

import "fmt"

type User struct {
	Name string
	Age  int
}

func main() {
	u := User{
		Name: "John",
		Age:  42,
	}
	fmt.Printf("%#v\n", u)
}

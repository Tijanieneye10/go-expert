package main

import (
	"fmt"
	"go-expert/oop"
)

func main() {
	person := &oop.User{
		Firstname: "John",
		Lastname:  "Doe",
		Age:       25,
	}

	fmt.Println(person.FullName())
}

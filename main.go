package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "John Doe"
	s2 := strings.Clone(s1)

	fmt.Println(s1, s2)
}

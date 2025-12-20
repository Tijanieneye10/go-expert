package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	s1 := "John Doe"
	s2 := strings.Clone(s1)

	s3 := strings.TrimSpace(s2)

	s4 := unicode.IsLetter(rune(s3[0]))

	fmt.Println(s4)

	value := fmt.Sprintf("%#v", s3)
	fmt.Println(value)
}

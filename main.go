package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "John Doe"
	s2 := strings.Clone(s1)

	s3 := strings.TrimSpace(s2)

	value := fmt.Sprintf("%+v", s3)
	fmt.Println(value)
}

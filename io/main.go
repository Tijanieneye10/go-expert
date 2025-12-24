package io

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := os.WriteFile("output.txt", []byte("Hello world, we just write to go file"), 0777)

	if err != nil {
		log.Fatal(err)
	}

	content, err := os.ReadFile("output.txt")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(string(content))
}

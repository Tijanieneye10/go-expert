package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	input := bufio.NewScanner(os.Stdin)

	fmt.Print("What is your name? ")

	if input.Scan() {
		name := input.Text()
		fmt.Println(name)
	}
}

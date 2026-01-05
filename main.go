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

		switch name {
		case "John":
			output := fmt.Sprintf("How are your? %s", name)
			fmt.Println(output)
		case "doe":
			output := fmt.Sprintf("Good morning %s", name)
			fmt.Println(output)
		default:
			fmt.Println("Invalid name")
		}

	}
}

package main

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string)

	go func() {
		fmt.Println("Hello World from main!")
		message <- "John Doe"
	}()

	time.Sleep(3 * time.Second)

	msg := <-message

	fmt.Println(msg)
}

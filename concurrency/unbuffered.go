package concurrency

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string) //Unbuffered channel

	go func() {
		fmt.Println("Hello World from main!")
		message <- "John Doe"
	}()

	time.Sleep(3 * time.Second)

	msg := <-message

	fmt.Println(msg)
}

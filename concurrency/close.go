package concurrency

import (
	"fmt"
	"time"
)

func main() {
	message := make(chan string, 3) //Buffered channel
	done := make(chan bool)

	go func() {
		fmt.Println("Hello World from main!")
		message <- "John Doe"
		message <- "Jane Doe"

		done <- true
	}()

	time.Sleep(3 * time.Second)

	fmt.Println(<-message)
	fmt.Println(<-message)

	m, ok := <-message
	if ok {
		fmt.Println(m)
	} else {
		fmt.Println("Channel closed")
		close(message)
	}
}

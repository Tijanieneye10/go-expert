package main

import (
	"fmt"
	"time"
)

func sayHello() {
	time.Sleep(2 * time.Second)
	fmt.Println("Hello World from SayHello!")
}

func main() {
	go sayHello()
	fmt.Println("Hello World from main!")
}

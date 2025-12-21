package main

import (
	"fmt"
	"sync"
	"time"
)

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(2 * time.Second)
	fmt.Println("Hello World from SayHello!")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(3)

	go sayHello(&wg)
	go sayHello(&wg)
	go sayHello(&wg)
	fmt.Println("Hello World from main!")

	wg.Wait()
}

package main

import (
	"fmt"
	"sync"
	"time"
)

/*
	* Rules
	1. Declare the wait outside
	2. Ensure the Add() increase equate the number of goroutine
	3.
*/

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

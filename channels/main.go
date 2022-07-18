package main

import (
	"fmt"
	"sync"
)

func main() {
	println("channels")

	ch := make(chan int, 5)
	wg := &sync.WaitGroup{}

	wg.Add(2)
	defer wg.Wait()
	go func(c <-chan int, wg *sync.WaitGroup) {
		// close(c)
		val, isChannelOpen := <-c
		fmt.Println(val, isChannelOpen)
		wg.Done()

	}(ch, wg)

	go func(c chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		// ch <- 6
		// close(ch)
		wg.Done()

	}(ch, wg)
}

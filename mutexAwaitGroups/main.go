package main

import (
	"fmt"
	"sync"
)

func main() {

	var score = []int{0}
	var wg = &sync.WaitGroup{}
	var mut = &sync.Mutex{}

	wg.Add(3)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("first")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()

	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("second")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()

	}(wg, mut)
	go func(wg *sync.WaitGroup, mut *sync.Mutex) {
		fmt.Println("third")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()

	}(wg, mut)
	wg.Wait()
	fmt.Println(score)
}

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // pointer
var signals = []string{"test"}
var mut sync.Mutex // pointer

func main() {

	websites := []string{"http://google.com", "http://facebook.com", "http://golang.org"}

	for _, site := range websites {
		go getStatusCode(site)
		wg.Add(1)
	}

	wg.Wait()
	fmt.Println(signals)

}

func getStatusCode(endPoint string) int {
	defer wg.Done()

	r, e := http.Get(endPoint)

	if e != nil {
		panic(e)
	}
	mut.Lock()
	signals = append(signals, endPoint)
	mut.Unlock()
	fmt.Println(r.StatusCode)
	return r.StatusCode
}

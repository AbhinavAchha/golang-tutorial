package main

import (
	"fmt"
	"net/http"
)

func main() {
    println("Hello, world!")
}

func greeter() {
    fmt.Println("Hello, world!")
}

func serveHome(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("<p>This is the home page.</p>"))
    
}

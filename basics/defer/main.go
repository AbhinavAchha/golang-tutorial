package main

import "fmt"

func main() {

    defer fmt.Println("World!")
    defer fmt.Println("Hello!1")
    defer fmt.Println("Hello!2")
    println("Hello, world!")
    deferFunc()
}

func deferFunc() {
    for i := 0; i < 10; i++ {
        defer fmt.Println(i)
    }
        
}

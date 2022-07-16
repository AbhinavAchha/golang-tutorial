package main

import "fmt"

func main() {
    println("Hello, world!")
    greeter()

    var result uint8 = adder(1, 2)
    result, a__ := proAddr(1, 2, 3, 4, 5)
    fmt.Println(result, a__)
}

func proAddr(vals ...uint8) (uint8, string) {
    var sum uint8
    for _, val := range vals {
        sum += val
    }
    return sum, "fasd"
}

func adder (a uint8, b uint8) uint8 {
    return a + b
}

func greeter() {
    println("Namaste, world!")
}

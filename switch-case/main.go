package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
    
    rand.Seed(time.Now().UnixNano())

    var diceNumber = rand.Intn(6) + 1

    fmt.Println(diceNumber)

    // if the current case is true then the fallthrough statement will execute
    // and the next case will be executed
    switch diceNumber {
    case 1:
        fmt.Println("You rolled a 1")
    case 2:
        fmt.Println("You rolled a 2")
    case 3:
        fmt.Println("You rolled a 3")
    case 4:
        fmt.Println("You rolled a 4")
        fallthrough
    case 5:
        fmt.Println("You rolled a 5")
    case 6:
        fmt.Println("You rolled a 6")
        fallthrough
    default:
        fmt.Println("You rolled a wrong number")
    }
}

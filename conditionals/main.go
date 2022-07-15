package main

import "fmt"

func main() {
    
    var count int = 23
    var result string

    if count < 10 {
        result = "less than 10"
    } else if count == 23 {
        result = "equal to 23"
    } else {
        result = "more than 10"
    }

    if true {
        fmt.Println("1 is true")
    }

    fmt.Println(result)

    if num := 5; num ==5 {
        fmt.Println("num is 5")
    } else {
        fmt.Println("num is not 5")
    }

    if 0 != 5 {
        fmt.Println("num is not 5")
    } else {
        fmt.Println("num is 5")
    }

}

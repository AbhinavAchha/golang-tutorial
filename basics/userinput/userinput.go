package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
    var wlcm string = "Welcome to Go!"

    fmt.Println("Enter your name: ")
    reader := bufio.NewReader(os.Stdin)

    name, _ := reader.ReadString('\n')
    fmt.Println(wlcm, name)
    fmt.Printf("%T", name)
}

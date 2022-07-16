package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
    println("Hello, world!")

    var content string = "Hello, world!"

    f, e  := os.Create("./hello.txt")

    if e != nil {
        panic(e)
    }

    l, e := io.WriteString(f, content)

    if e != nil {
        panic(e)
    }

    fmt.Println(l)
    defer f.Close()

    readFile("./hello.txt")
}

func readFile(fileName string) {
    var f, e = ioutil.ReadFile(fileName)

    if e != nil {
        panic(e)
    }

    fmt.Println(string(f))
}

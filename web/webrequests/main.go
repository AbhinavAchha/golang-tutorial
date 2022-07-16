package main

import (
	"fmt"
	"io"
	"net/http"
)

const url = "https://lco.dev"

func main() {
    var r, e =  http.Get(url)

    if e != nil {
        panic(e)
    } 

    fmt.Printf("%T\n", r)
    fmt.Printf("%+v\n", r)
    defer r.Body.Close() // must close Body when done reading

    var data, er = io.ReadAll(r.Body)
    if er != nil {
        panic(er)
    }
    var content = string(data)
    fmt.Println(content)
}

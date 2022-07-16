package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
    // GetRequest("https://www.google.com")
    PostJsonRequest("http://localhost:8000/post")
    
}

func GetRequest(url string) {
    var response, err = http.Get(url)

    if err != nil {
        panic(err)
    }

    defer response.Body.Close()

    fmt.Println(response.StatusCode)
    fmt.Println(response.ContentLength)

    var responseString strings.Builder

    var content, _ = io.ReadAll(response.Body)
    byteCount, _ := responseString.Write(content)
    fmt.Println(byteCount)
    fmt.Println(responseString.String())
    // fmt.Println(string(content))
}

func PostJsonRequest(url string)  {

    var requestBody = strings.NewReader(`
    {
        "courseName": "Go Web Development",
        "courseDescription": "Go Web Development Course",
        "coursePrice": "100"
    }
    `)

    var response, err = http.Post(url, "application/json", requestBody)

    if err != nil {
        panic(err)
    }

    defer response.Body.Close()

    var content, e = io.ReadAll(response.Body)

    if e != nil {
        panic(e)
    }

    fmt.Println(string(content))
    
}

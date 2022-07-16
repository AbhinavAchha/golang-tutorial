package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
    Name string `json:"courseName"`
    Price int64 `json:"price"`
    password string `json:"-"`
    Tags []string `json:"tags,omitempty"`
}

func main() {
    
    EncodeJson()
}

func EncodeJson() {

    var courses = []course{
        {"Golang", 100, "123", []string{"go", "backend"}},
        {"Python", 200, "456", []string{"python", "backend"}},
        {"React", 300, "891", nil},
    }

    var jsonData, e = json.MarshalIndent(courses, "", "\t")
    if e != nil {
        panic(e)
    }
    fmt.Printf("%s\n", jsonData)
}

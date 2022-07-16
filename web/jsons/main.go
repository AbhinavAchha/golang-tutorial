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
    
    DecodeJson()
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

func DecodeJson()  {
   var jsonData = []byte(`
       {
                "courseName": "Golang",
                "price": 100,
                "tags": ["go","backend"]
        }
   `) 

   var cource course

   var checkValid = json.Valid(jsonData)
    if checkValid {
        fmt.Println("Valid json")
        json.Unmarshal(jsonData, &cource)
        fmt.Printf("%#v\n", cource)
    } else {
        fmt.Println("Invalid json")
    }

    var onlineData map[string]interface{}
    json.Unmarshal(jsonData, &onlineData)
    fmt.Printf("%#v\n", onlineData)

    for k, v := range onlineData {
        fmt.Printf("%s: %v\n", k, v)
    }
}

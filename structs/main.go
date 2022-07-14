package main

import "fmt"

func main() {

    new_user := User{"user", "user@123", true, 1}

    new_user.Age = 2

    fmt.Println(new_user)
    fmt.Printf("%+v\n", new_user)
    
}
    type User struct {
        Name string
        Email string
        Status bool
        Age int8
    }

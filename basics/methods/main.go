package main

import "fmt"

func main() {

    var user =  User{Name: "Name", Email: "EMail", Alive: true}

    fmt.Println(user.GetAlive())
    user.NewMail()
    fmt.Println(user.Email)
    
}

type User struct {
    Name string
    Age int
    Alive bool
    Email string
}

func (u User) GetAlive() bool{
    fmt.Println(u.Alive)
    return u.Alive
}

// NewMail is a method of the User struct
// This method takes a `copy` of the User struct and changes the Email field
// Thus the original User struct is not changed
func (u User) NewMail() {
    u.Email = "newmail@mail.com"
    fmt.Println(u.Email)
}

package main

import "fmt"

const Token = "token";

func main() {
    var username string = "user";
    var isUser bool = true;
    fmt.Printf("Hello %T\n", username);
    fmt.Printf("Hello %T\n", isUser);


    var small uint = 255;
    fmt.Printf("%T\n", small);

    var float float32 = 1.2345;
    fmt.Printf("%T\n", float);

    var i string;
    fmt.Printf("%T\n", i);
    fmt.Println("jfas", i, "dsjfh")

    numberofUsers := 10;
    fmt.Printf("%T\n", numberofUsers);
    fmt.Println(numberofUsers)

}

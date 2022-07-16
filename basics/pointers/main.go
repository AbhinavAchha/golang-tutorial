package main;

import "fmt";

func main() {
	fmt.Println('b');
    
    // var n *int
    // fmt.Println(n);

    n := 33

    var ptr = &n;

    fmt.Println(ptr);
    fmt.Println(*ptr);
};

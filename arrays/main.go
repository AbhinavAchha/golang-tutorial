package main

import "fmt"

func main() {

	var list [10]string

	list[0] = "a"
	list[1] = "b"
	// list[2] = "c";
	list[3] = "d"

	fmt.Println(list, len(list))

    var list2 =  [4]string{"a", "b", "c", "d"};
    fmt.Println(list2, len(list2))

}

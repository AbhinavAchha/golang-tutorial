package main

import (
	"fmt"
	"sort"
)

func main() {
	var list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	fmt.Printf("%T\n", list)

    list = append(list, 11)
    fmt.Println(list)

    list = append(list[1:])

    fmt.Println(list)

    highScore := make([]int, 4)

    highScore[0] = 100
    highScore[1] = 13
    highScore[2] = 89
    highScore[3] = 95
    // highScore[4] = 56


    highScore = append(highScore, 56, 8121, 19283)
    sort.Ints(highScore)
    fmt.Println(highScore)

}

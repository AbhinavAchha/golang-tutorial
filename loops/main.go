package main

import "fmt"

func main() {
    
    // var days = []string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday", "Sunday"}

    // for d:=0; d<len(days); d++ {
    //     fmt.Println(days[d])
    // }
    //
    // for i := range days {
    //     fmt.Println(days[i])
    // }
    //
    // for i, d := range days {
    //     fmt.Println(i, d)
    // }

    var v = 0

    for v < 10 {
        v++

        if v == 2 {
            goto label 
        }
        
        if v ==5 {
            continue
        }
        fmt.Println(v)
    }

    label: fmt.Println("label")
}

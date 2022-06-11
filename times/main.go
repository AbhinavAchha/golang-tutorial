package main

import (
	"fmt"
	"time"
)

func main() {
    fmt.Println("ok")

    currentTime := time.Now()
    fmt.Println(currentTime.Format("2006-01-92 Monday 15:04:05"))

    newDate := time.Date(2021, time.April, 10, 12, 44, 44, 44, time.UTC)

    fmt.Println(newDate.Format("2006-01-02 Monday 15:04:05"))

}

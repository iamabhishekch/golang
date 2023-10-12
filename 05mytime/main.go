package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("welcome to the study of time")
	presentTime := time.Now()
	fmt.Println(presentTime)

	fmt.Println(presentTime.Format("01-02-2006 15:04:05 Monday")) // we always have to use the same format like for day we have to write Monday

	createdTime := time.Date(2023, time.October, 24, 14, 59, 48, 38, time.UTC) // only month will be written in "time.Month" format
	fmt.Println(createdTime)

	fmt.Println(createdTime.Format("01-02-2006 15:04:05 Monday"))  
}

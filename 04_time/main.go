package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time in Golang")
	t := time.Now()
	fmt.Println(t)
	fmt.Println(t.Format("01-02-2006 15:04:05 Monday"))
}
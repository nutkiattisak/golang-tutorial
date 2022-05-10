package main

import (
	"fmt"
	"time"
)

func process1(ch chan string, data string) {
	ch <- data
}

func main() {
	ch := make(chan string, 1)
	go process1(ch, "Data")
	fmt.Println(<-ch)
	time.Sleep(5 * time.Second)
}
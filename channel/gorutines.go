package main

import (
	"fmt"
	"time"
)

func f(from string) {
	for i := 0; i < 1003; i++ {
		fmt.Println(from, ":", i)
	}
}

func main() {
	go f("Hello")
	go f("Hello 2")
	time.Sleep(5 * time.Second)
}
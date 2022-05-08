package main

import "fmt"

const count = 10

func main() {
	index := 0
	for index < 10 {
		fmt.Println(index)
		index++
	}

	for i := 0; i < count; i++ {
		fmt.Println(i)
	}

	var input string
	for {
		fmt.Println("loop")
		fmt.Scanf("%s", &input)
		if input == "exit" {
			break
		}

	}
}
package main

import "fmt"

func hello() {
	fmt.Println("Hello World")
}

func plus(number1 int, number2 int) int {
	return number1 + number2
}

func main() {
	hello()
	result := plus(5, 4)
	fmt.Println(result)
}
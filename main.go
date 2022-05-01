package main

import "fmt"

var numberInt, numberInt2 int  = 1000, 2000
var message string = "Hello "

func main() {
	numberfloat := 25.4
	fmt.Println(numberInt)
	fmt.Println(numberInt2)
	fmt.Println(numberfloat)
	fmt.Println(message)
	fmt.Println(float64(numberInt) + numberfloat)
	fmt.Println(message + "Kiattisak")
	fmt.Println("ข้อความกับตัวเลข", numberInt)
}
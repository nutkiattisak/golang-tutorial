package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var reader = bufio.NewReader(os.Stdin)

func getInput(promt string) float64 {
	fmt.Printf("%v",promt)
	input, _ := reader.ReadString('\n')
	value, error := strconv.ParseFloat(strings.TrimSpace(input), 64) //แปลงค่าเป็นทศนิยม
	if error != nil {
		message, _ := fmt.Scanf("%v must number only", &promt)
		panic(message) //แสดงข้อความคล้าย print
	}

	return value
}

func add(value1, value2 float64) float64 {
	return value1 + value2
}

func subtract(value1, value2 float64) float64 {
	return value1 - value2
}

func multiply(value1, value2 float64) float64 {
	return value1 * value2
}

func divide(value1, value2 float64) float64 {
	return value1 / value2
}

func getOperator() string {
	fmt.Print("Operator is (+ - * /) = ")
	op, _ := reader.ReadString('\n')
	return strings.TrimSpace(op)
}

func main() {
	value1 := getInput("value 1 = ")
	value2 := getInput("value 2 = ")

	var result float64

	switch operator := getOperator(); operator {
		case "+":
			result = add(value1, value2)
		case "-":
			result = subtract(value1, value2)
		case "*":
			result = multiply(value1, value2)
		case "/":
			result = divide(value1, value2)
		default:
			panic("Wrong operator")
	}

	fmt.Printf("Result is %v", result)
}
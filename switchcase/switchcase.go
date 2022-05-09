package main

import (
	"fmt"
)

var input string
var colorCode string

func main() {
	fmt.Scan(&input)
	switch input {
		case "blue":
			fmt.Println("#0000FF")
		case "red":
			fmt.Println("#FF0000")
		case "green":
			fmt.Println("#00FF00")
		case "yellow":
			fmt.Println("#FFFF00")
		default:
			fmt.Println("unknown")
	}
}
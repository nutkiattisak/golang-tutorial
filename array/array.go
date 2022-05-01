package main //pk+tap

import "fmt"

var productName [4]string

func main() {
	fmt.Println(productName)

switch {
	case productName[0] == "":
		fmt.Println("one")
	case productName[1] == "":
		fmt.Println("two")
	case productName[2] == "":
		fmt.Println("three")
	}
}
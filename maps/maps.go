package main

import "fmt"

var product = make(map[string]float64)

func main() {
	fmt.Println(product)

	//Create
	product["A"] = 1.0
	product["B"] = 2.0
	product["C"] = 3.0
	product["D"] = 4.0
	fmt.Println(product)

	//Delete
	delete(product, "A")
	fmt.Println(product)

	//Update
	product["B"] = 5.0
	fmt.Println(product)

	value1 := product["B"]
	fmt.Println(value1)

	value2 := map[string]float64{"A": 1.0, "B": 2.0, "C": 3.0, "D": 4.0}
	fmt.Println(value2)
	
}
package main

import "fmt"

func add(a, b int) {
	result := a + b
	fmt.Println("result:", result)
}

func loop() {
	for i := 0; i < 10; i++ {
		defer add(i, i)
	}
}

func main() {
	defer fmt.Println("Goodbye!")
	fmt.Println("Welcome to the playground!")
	defer add(1, 2)
	defer add(10, 20) 
	loop() // แสดงลำดับแรก
	//LIFO Last in first out
}
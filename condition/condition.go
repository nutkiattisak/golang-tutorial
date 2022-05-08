package main

import "fmt"

func main() {

	var input int
	myMoney := 100
	if myMoney > 100 {
		fmt.Println("I have more than 100")
	} else if myMoney == 100 {
		fmt.Println("I have 100")
	} else {
		fmt.Println("I have less than 100")
	}

	fmt.Scanf("%d", &input)
	
	if input == 1 {
		fmt.Println("You have entered 1")
	} else {
		fmt.Println("You have entered something else")
	}

}
package main

import "fmt"

type employee struct {
	employeeID   string
	employeeName string
	phone        string
}

type product struct {
	productID   string
	productName string
	price       float64
}

func main() {
	//แบบจอง
	employeeList := [3]employee{}
	employeeList[0] = employee {
		employeeID:   "101",
		employeeName: "John",
		phone:        "0123456789",
	}
	employeeList[1] = employee {
		employeeID:   "102",
		employeeName: "Ami",
		phone:        "0876543210",
	}
	employeeList[2] = employee {
		employeeID:   "103",
		employeeName: "Raj",
		phone:        "0412345678",
	}

	//แบบไม่จอง
	productList := []product{}
	product1 := product {
		productID:   "101",
		productName: "Laptop",
		price:       100,
	}

	productList = append(productList, product1)

	fmt.Println("Employee List: ", employeeList)
	fmt.Println("Product List", productList)

	
}
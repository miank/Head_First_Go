package main

import (
	"fmt"
)

func main() {
	var quantity int
	var length, width float64
	var customerName string
	quantity = 2
	customerName = "Damon Cole"
	length, width = 1.2, 2.4
	fmt.Println("Quantity ", quantity)
	fmt.Println("Length and Width", length, width)
	fmt.Println("Customer Name ", customerName)
	// Zero values
	var myInt int
	var myFloat float64
	fmt.Println("Initial Values ", myInt, myFloat)
	// Short variable decleration
	quantity1 := 4
	var length1, width1 float64 = 1.2, 2.4
	fmt.Println("Quantity1 ", quantity1)
	fmt.Println("length1 and width1 ", length1, width1)
}

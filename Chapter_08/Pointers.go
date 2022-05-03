package main

import "fmt"

type myStruct struct {
	myField int
}

func main() {
	var value int = 2
	var pointer *int = &value
	fmt.Println("Address ", pointer)
	fmt.Println("Value ", *pointer)

	var value1 myStruct
	value1.myField = 3
	var pointer1 *myStruct = &value1 // creating a pointer variable and assign the address of the structure.
	pointer1.myField = 9
	fmt.Println((*pointer1).myField)
	fmt.Println(pointer1.myField)

}

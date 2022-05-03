package main

import "fmt"

var pet struct {
	name string
	age  int
}

func main() {
	pet.name = "Max"
	pet.age = 5
	fmt.Println("Name :", pet.name)
	fmt.Println("Age :", pet.age)
}

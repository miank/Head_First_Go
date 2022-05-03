package main

import "fmt"

type student struct {
	name  string
	grade float64
}

func main() {
	var s student
	s.name = "Alonzo Cole"
	s.grade = 92.3
	fmt.Println("Name : ", s.name)
	fmt.Println("Garde : ", s.grade)
}

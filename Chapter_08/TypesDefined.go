package main

import "fmt"

type part struct {
	description string
	count       int
}

type car struct {
	name     string
	topSpeed int
}

func main() {
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Printf("%#v\n", porsche)
	fmt.Println("Porsche Name :", porsche.name)
	fmt.Println("Top Speed :", porsche.topSpeed)
	var bolts part
	bolts.description = "Hex Bolts"
	bolts.count = 24
	fmt.Printf("%#v\n", bolts)
	fmt.Println("Description :", bolts.description)
	fmt.Println("Count :", bolts.count)

}

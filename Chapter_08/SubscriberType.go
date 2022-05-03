package main

import "fmt"

type Subscriber struct {
	name   string
	rate   float64
	active bool
}

func printInfo(s Subscriber) {
	fmt.Println("Name : ", s.name)
	fmt.Println("Monthly rate : ", s.rate)
	fmt.Println("Active?", s.active)
}

func defaultSubscriber(name string) Subscriber {
	var s Subscriber
	s.name = name
	s.rate = 5.99
	s.active = true
	return s
}
func applyDiscount(s *Subscriber) {
	s.rate = 4.99
}

func main() {
	subscriber1 := defaultSubscriber("Aman Singh")
	subscriber1.rate = 4.99
	printInfo(subscriber1)

	subscriber2 := defaultSubscriber("Beth Ryan")
	printInfo(subscriber2)

	var s Subscriber
	fmt.Println(s.rate) // Before
	applyDiscount(&s)
	fmt.Println(s.rate) // After
}

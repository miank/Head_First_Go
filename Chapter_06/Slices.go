package main

import "fmt"

func main() {

	var notes []string
	notes = make([]string, 7)
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	primes := make([]int, 5)
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])
	fmt.Println(len(notes))
	fmt.Println(len(primes))

	letters := []string{"a", "b", "c"}
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}
	for _, letter := range letters {
		fmt.Println(letter)
	}

	notes = []string{"do", "re", "mi", "fa", "so", "la", "ti"}
	fmt.Println(notes[3], notes[6], notes[0])
	primes = []int{
		2,
		3,
		5,
	}
	fmt.Println(primes[0], primes[1], primes[2])

	// Slice operator
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[0:3]
	fmt.Println(slice1)

	i, j := 1, 4
	slice2 := underlyingArray[i:j]
	fmt.Println(slice2)
	slice3 := underlyingArray[2:5]
	fmt.Println(slice3)
	slice4 := underlyingArray[:3]
	fmt.Println(slice4)
	slice5 := underlyingArray[1:]
	fmt.Println(slice5)
}

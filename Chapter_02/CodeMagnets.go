package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("my.txt") // create this file before running the code
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.Size())

}

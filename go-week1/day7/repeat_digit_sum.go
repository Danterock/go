package main

import "fmt"

func main() {
	var number int
	var i int
	fmt.Print("Enter number: ")
	fmt.Scan(&number)
	for number != 0 {
		i += number % 10
		number = number / 10

	}
	fmt.Printf("Digit sum: %v", i)
}

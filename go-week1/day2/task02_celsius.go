package main

import "fmt"

func main() {
	var C float64
	fmt.Print("Enter a number: ")
	fmt.Scan(&C)
	fmt.Println("Fahrenheit:", C*9/5+32)
}

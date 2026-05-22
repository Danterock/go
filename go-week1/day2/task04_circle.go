package main

import "fmt"

func main() {
	const pi = 3.14159
	var radius float64
	fmt.Print("Enter radius of circle: ")
	fmt.Scan(&radius)
	fmt.Printf("Area: %f\nLength: %f", pi*radius*radius, 2*pi*radius)
}

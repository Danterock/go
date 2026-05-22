package main

import "fmt"

func main() {
	var radius float64
	fmt.Print("Enter radius of circle: ")
	fmt.Scan(&radius)
	fmt.Printf("Area: %f\nLength: %f", circleArea(radius), circleLength(radius))
}

func circleArea(radius float64) float64 {
	const pi = 3.14159
	return pi * radius * radius
}

func circleLength(radius float64) float64 {
	const pi = 3.14159
	return 2 * pi * radius
}

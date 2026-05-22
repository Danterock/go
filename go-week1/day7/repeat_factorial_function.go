package main

import "fmt"

func main() {
	var fac int
	fmt.Printf("Enter number: ")
	fmt.Scan(&fac)
	if fac < 0 {
		fmt.Println("Invalid number")
	} else if fac == 0 {
		fmt.Println("Factorial: 1")
	} else {
		fmt.Printf("Factorial: %v", factorial(fac))
	}
}

func factorial(n int) int {
	var result int = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}

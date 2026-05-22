package main

import "fmt"

func main() {
	var f int
	fmt.Print("Input number ")
	fmt.Scan(&f)
	if f < 0 {
		fmt.Println("invalid number")
	} else {
		fmt.Println("Factorial:", factorial(f))
	}
}

func factorial(n int) int {
	var number int = 1
	for i := 1; i <= n; i++ {
		number *= i
	}
	return number
}

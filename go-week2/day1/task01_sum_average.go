package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Print("Input capacity of slice: ")
	fmt.Scan(&n)
	if n < 1 {
		fmt.Println("Invalid n")
	}
	numbers := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Printf("input number %v\n", i+1)
		fmt.Scan(&numbers[i])
	}
	fmt.Printf("Average: %v\n", float64(average(numbers)))
	fmt.Printf("Sum: %v\n", int(sum(numbers)))
}

func sum(numbers []int) int {
	num := 0
	for i := 0; i < len(numbers); i++ {
		num += numbers[i]
	}
	return num
}
func average(numbers []int) float64 {
	var sum float64
	for i := 0; i < len(numbers); i++ {
		sum += float64(numbers[i])
	}
	return sum / float64(len(numbers))
}

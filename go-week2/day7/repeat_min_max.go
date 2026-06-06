package main

import "fmt"

func main() {
	var n int

	fmt.Println("Input number")
	fmt.Scan(&n)
	var numbers = make([]int, n)
	if n < 1 {
		fmt.Println("invalid number")
	}
	for i := 0; i < n; i++ {
		fmt.Printf("Input %v number: ", i+1)
		fmt.Scan(&numbers[i])
	}
	min, max := minMax(numbers)
	fmt.Printf("Min: %v Max: %v\n", min, max)
}

func minMax(numbers []int) (int, int) {
	var min, max int = 1, 1
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] > max {
			max = numbers[i]
		}
	}
	return min, max
}

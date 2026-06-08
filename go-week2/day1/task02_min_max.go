package main

import "fmt"

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
	min, max := minMax(numbers)
	fmt.Printf("Min: %v\nMax: %v\n", min, max)
}

func minMax(numbers []int) (int, int) {
	if len(numbers) == 0 {
		return 0, 0
	}
	if len(numbers) == 1 {
		return numbers[0], numbers[0]
	}
	var min, max int = numbers[0], numbers[0]
	for i := 0; i < len(numbers); i++ {
		if numbers[i] < min {
			min = numbers[i]
		}
		if numbers[i] >= max {
			max = numbers[i]
		}
	}
	return min, max
}

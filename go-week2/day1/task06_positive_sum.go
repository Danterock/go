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
	a := sumPositive(numbers)
	if a == 0 {
		fmt.Println("Sum: 0")
	} else {
		fmt.Printf("Sum: %v\n", a)
	}

}

func sumPositive(numbers []int) int {
	var sum int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] > 0 {
			sum += numbers[i]
		}
	}
	return sum
}

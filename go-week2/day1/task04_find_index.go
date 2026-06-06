package main

import "fmt"

func main() {
	var n, target int
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
	fmt.Print("Input target: ")
	fmt.Scan(&target)
	a := findIndex(numbers, target)
	if a == -1 {
		fmt.Println("Not found")
	} else {
		fmt.Printf("Index: %v\n", a)
	}
}

func findIndex(numbers []int, target int) int {
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == target {
			return i
		}
	}
	return -1
}

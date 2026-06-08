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
	fmt.Printf("Even numbers: %v\n", filterEven(numbers))

}

func filterEven(numbers []int) []int {
	var newnumbers []int
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			newnumbers = append(newnumbers, numbers[i])
		}
	}
	return newnumbers
}

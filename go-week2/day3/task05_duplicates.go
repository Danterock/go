package main

import (
	"fmt"
	"sort"
)

func main() {
	var n int
	fmt.Print("Input quantity: ")
	fmt.Scan(&n)

	if n < 1 {
		fmt.Println("Invalid n")
		return
	}

	numbers := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Println("Input number: ", i+1)
		fmt.Scan(&numbers[i])
	}

	duplicates := findDuplicates(numbers)

	if len(duplicates) == 0 {
		fmt.Println("No duplicates")
		return
	}

	fmt.Print("Duplicates:")

	for i := 0; i < len(duplicates); i++ {
		fmt.Print(" ", duplicates[i])
	}
}

func findDuplicates(numbers []int) []int {
	freq := make(map[int]int)

	for i := 0; i < len(numbers); i++ {
		freq[numbers[i]]++
	}

	duplicates := []int{}

	for number, count := range freq {
		if count > 1 {
			duplicates = append(duplicates, number)
		}
	}

	sort.Ints(duplicates)

	return duplicates
}

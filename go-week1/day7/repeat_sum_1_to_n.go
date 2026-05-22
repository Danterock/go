package main

import "fmt"

func main() {
	var sum, res int
	fmt.Printf("Enter a number: ")
	fmt.Scan(&sum)
	if sum < 0 {
		fmt.Println("error: sum is negative")
	} else {
		for i := 1; i <= sum; i++ {
			res += i
		}
		fmt.Printf("Sum: %v\n", res)
	}
}

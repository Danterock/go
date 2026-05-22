package main

import "fmt"

func main() {
	var a, b, c int
	fmt.Print("Input three numbers")
	fmt.Scan(&a, &b, &c)
	fmt.Println(max3(a, b, c))
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func max3(a int, b int, c int) int {
	if max(a, b) > c {
		return max(a, b)
	}
	return c
}

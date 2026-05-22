package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter number: ")
	fmt.Scan(&n)
	a, b := countEvenOdd(n)
	fmt.Printf("Even: %v\nOdd: %v\n", b, a)
}

func countEvenOdd(n int) (int, int) {
	var a int
	var b int
	if n < 1 {
		return 0, 0
	}
	for i := 0; i < n; i++ {
		if i%2 == 0 {
			a += 1
		} else {
			b += 1
		}
	}
	return a, b
}

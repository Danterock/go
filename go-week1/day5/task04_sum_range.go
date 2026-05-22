package main

import "fmt"

func main() {
	var a, b int
	fmt.Print("Input two numbers: ")
	fmt.Scan(&a, &b)
	fmt.Println("Sum:", sumRange(a, b))
}

func sumRange(a, b int) int {
	n := 0
	if a > b {
		for i := 0; i <= a; i++ {
			if b <= i {
				n += i
			}
		}
	} else {
		for i := 0; i <= b; i++ {
			if a <= i {
				fmt.Println(i)
				n += i
			}
		}
	}
	return n
}

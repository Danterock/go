package main

import "fmt"

func main() {
	var a int
	var b int
	fmt.Println("Enter two numbera:")
	fmt.Scan(&a, &b)
	if a > b {
		fmt.Println("Max:", a)
	} else {
		fmt.Println("Max:", b)
	}
}

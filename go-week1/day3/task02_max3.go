package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Println("Enter three numbers:")
	fmt.Scan(&a, &b, &c)
	if a >= b && a >= c {
		fmt.Println("Max:", a)
	} else if b >= c && b >= a {
		fmt.Println("Max:", b)
	} else if c >= b && c >= a {
		fmt.Println("Max:", c)
	} else {
		fmt.Println("error")
	}
}

package main

import "fmt"

func main() {
	var a, b int
	fmt.Println("Input two numbers")
	fmt.Scan(&a, &b)
	calc(a, b)
}

func calc(a, b int) {
	fmt.Println("Sum:", a+b)
	fmt.Println("Difference:", a-b)
	fmt.Println("Product:", a*b)
	fmt.Println("Division:", float64(a)/float64(b))
}

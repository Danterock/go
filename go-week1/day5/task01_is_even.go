package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter a number: ")
	fmt.Scan(&n)
	if isEven(n) == true {
		fmt.Println("Even")
	} else {
		fmt.Println("Odd")
	}

}

func isEven(n int) bool {
	if n%2 == 0 {
		return true
	}
	return false
}

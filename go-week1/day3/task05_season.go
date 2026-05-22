package main

import "fmt"

func main() {
	var n int
	fmt.Print("Enter your number: ")
	fmt.Scan(&n)
	if n == 12 || n == 1 || n == 2 {
		fmt.Println("Winter")
	} else if n >= 3 && n <= 5 {
		fmt.Println("Spring")
	} else if n >= 6 && n <= 8 {
		fmt.Println("Summer")
	} else if n >= 9 && n <= 11 {
		fmt.Println("Autumn")
	} else {
		fmt.Println("Invalid month")
	}
}

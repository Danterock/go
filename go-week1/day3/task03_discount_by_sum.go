package main

import "fmt"

func main() {
	var amount int
	fmt.Print("Enter amount:")
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Invalid amount")
	} else if 0 < amount && amount <= 999 {
		fmt.Printf("Discount: 0\nFinal price: %v", amount)
	} else if 999 < amount && amount <= 4999 {
		fmt.Printf("Discount: 5\nFinal price: %f", float64(amount)-float64(amount)*0.05)
	} else {
		fmt.Printf("Discount: 10\nFinal price: %f", float64(amount)-float64(amount)*0.1)
	}

}

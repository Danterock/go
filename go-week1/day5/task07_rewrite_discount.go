package main

import "fmt"

func main() {
	var amount float64
	fmt.Print("Enter amount:")
	fmt.Scan(&amount)
	if amount <= 0 {
		fmt.Println("Invalid amount")
	} else {
		fmt.Printf("Discount: %v\nFinal price: %v", getDiscountPercent(amount), getFinalPrice(amount, getDiscountPercent(amount)))
	}

}

func getDiscountPercent(amount float64) int {
	if 0 < amount && amount <= 999 {
		return 0
	} else if 999 < amount && amount <= 4999 {
		return 5
	}
	return 10
}

func getFinalPrice(amount float64, discount int) float64 {
	return amount - amount*float64(discount)/100

}

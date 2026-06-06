package main

import (
	"fmt"
)

func main() {
	stock := map[string]int{
		"apple": 10,
		"milk":  5,
		"bread": 3,
	}
	var product string
	var quantity int
	fmt.Print("Input product name and quantity: ")
	fmt.Scanln(&product, &quantity)
	if quantity < 1 {
		fmt.Println("invalid quantity")
		return
	}
	_, ok := stock[product]
	if !ok {
		fmt.Println("Unknown product")
		return
	}
	if !buyProduct(stock, product, quantity) {
		fmt.Println("Not enough stock")
		return
	}
	fmt.Println("Remaining: ", stock[product])
}

func buyProduct(stock map[string]int, product string, quantity int) bool {
	if stock[product] < quantity {
		return false
	}
	stock[product] -= quantity
	return true
}

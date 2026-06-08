package main

import "fmt"

type Product struct {
	Title    string
	Price    float64
	Discount int
}

func (p Product) DiscountedPrice() float64 {
	if p.Price < 0 || p.Discount < 0 || p.Discount > 100 {
		fmt.Println("Invalid product")
		return 0.0
	}
	return p.Price - p.Price*float64(p.Discount)/100
}

func main() {
	var product Product
	var result float64
	fmt.Println("Input all information:")
	fmt.Scan(&product.Title)
	fmt.Scan(&product.Price)
	fmt.Scan(&product.Discount)
	result = product.DiscountedPrice()
	fmt.Println("Final price:", result)
}

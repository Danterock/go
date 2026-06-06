package main

import "fmt"

type Product struct {
	ID      int
	Title   string
	Price   float64
	InStock bool
}

var maxPrice float64
var products = []Product{
	{1, "Mouse", 1200, true},
	{2, "Keyboard", 3000, true},
	{3, "Monitor", 12500, false},
	{4, "Cable", 450, true},
}

func filterProducts(products []Product, maxPrice float64) []Product {
	res := []Product{}
	for i := 0; i < len(products); i++ {
		if products[i].InStock && products[i].Price <= maxPrice {
			res = append(res, products[i])
		}
	}
	return res
}

func print(maxPrice float64) []string {
	var res []string
	sorted := filterProducts(products, maxPrice)

	if len(sorted) == 0 {
		res = append(res, "No products")
		return res
	}

	for i := 0; i < len(sorted); i++ {
		res = append(res, fmt.Sprintf("%d %s %.2f\n", sorted[i].ID, sorted[i].Title, sorted[i].Price))
	}
	return res
}

func main() {
	fmt.Printf("Input max price: ")
	fmt.Scan(&maxPrice)
	fmt.Println(print(maxPrice))
}

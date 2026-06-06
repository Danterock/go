package main

import "fmt"

type Product struct {
	ID      int
	Title   string
	Price   float64
	InStock bool
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

func main() {
	products := []Product{
		{1, "Mouse", 1200, true},
		{2, "Keyboard", 3000, true},
		{3, "Monitor", 12500, false},
		{4, "Cable", 450, true},
	}

	var maxPrice float64
	fmt.Printf("Input max price: ")
	fmt.Scan(&maxPrice)

	sorted := filterProducts(products, maxPrice)

	if len(sorted) == 0 {
		fmt.Println("No products")
	}

	for i := 0; i < len(sorted); i++ {
		fmt.Printf("%d %s %.2f\n", sorted[i].ID, sorted[i].Title, sorted[i].Price)
	}

}

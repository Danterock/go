package main

import "fmt"

func main() {
	var price float64
	var discount float64
	var finalprice float64
	fmt.Println("Input price and discount: ")
	fmt.Scan(&price, &discount)
	finalprice = price - price*discount/100
	fmt.Printf("Final price: %f", finalprice)
}

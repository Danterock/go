package main

import "fmt"

func main() {
	var n int
	var num int = 1
	fmt.Print("Enter number: ")
	fmt.Scan(&n)
	if n < 0 {
		fmt.Println("Invalid n")
	} else if n == 0 {
		fmt.Println("Factorial: 1")
	} else {
		for i := 1; i <= n; i++ {
			num *= i
		}
		fmt.Println("Fuctorial:", num)
	}
}

/*
количество прогонов цикла зависит от вводимого числа
возьмем для примера n = 3
при первом прогоне цикла мы записываем в переменную num значение i равное 1
при втором прогоне цикла мы записываем в переменную num значение num * i (1 * 2)
при третьем прогоне цикла мы записываем в переменную num значение num * i (2 * 3)
после чего цикл заканчивается и мы выводим факториал и num
*/

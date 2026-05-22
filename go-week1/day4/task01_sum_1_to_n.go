package main

import "fmt"

func main() {
	var n int
	var num int
	fmt.Print("Enter number: ")
	fmt.Scan(&n)
	if n < 1 {
		fmt.Println("Invalid n")
	} else {
		for i := 1; i <= n; i++ {
			num += i
		}
		fmt.Println("Sum:", num)
	}
}

/*
количество прогонов цикла зависит от вводимого числа
возьмем для примера n = 3
при первом прогоне цикла мы записываем в переменную num значение i равное 1
при втором прогоне цикла мы записываем в переменную num значение num + i (1 + 2)
при третьем прогоне цикла мы записываем в переменную num значение num + i (3 + 3)
после чего цикл заканчивается и мы выводим сумму и num
*/

package main

import "fmt"

func main() {
	var secret int = 37
	var score int = 1
	var n int
	for {
		fmt.Print("Enter your number: ")
		fmt.Scan(&n)
		if secret == n {
			fmt.Printf("Correct! Attempts: %v\n", score)
			break
		} else if secret > n {
			fmt.Println("greater")
		} else if secret < n {
			fmt.Println("less")
		}
		score++
	}
}

/*
тут у нас бесконечный цикл
который может быть завершен только при условии, что пользователь введет число 37
если пользователь вводит число,которое больше 37, срабатывает условие, результатом которого является вывод в консоль "много"
если пользователь вводит число меньше 37 - мы выводим сообщение "мало"
так же при каждом прогоне цикла есть переменная score которая является счетчиком и она увеличивается на 1
*/

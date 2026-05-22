package main

import (
	"fmt"
	"strconv"
)

func main() {
	var number int
	var result int
	fmt.Print("Enter number: ")
	fmt.Scan(&number)
	if number < 0 {
		fmt.Println("Invalid number")
	} else {
		str := strconv.Itoa(number)
		fmt.Println(str)
		for i := 0; i < len(str); i++ {
			result += number % 10
			number /= 10
		}
		fmt.Println(result)
	}
}

/*
количество прогонов цикла зависит от длины числа, которое ввел пользователь
для примера введем 1111
при первом прогоне в переменную result мы записываем последнюю цифру из введенного числа после чего мы делим число на 10
при втором прогоне в переменную result к 1 мы добавляем опять последнюю цифру из number там осталось число 111 и после этого опять делем его на 10
при третьем прогоне в переменную result к 2 мы опять добавляем последнюю цифру из number и делим на 10
при четвертом и последнем прогоне в переменную result к 3 мы добавляем 1, последнюю цифру из number
после этого цикл завершен
*/

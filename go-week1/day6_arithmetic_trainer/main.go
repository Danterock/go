package main

import "fmt"

var a, b, userAnswer, cor, wro int
var number int = 1

func main() {
	printMenu()
	for {
		fmt.Scan(&number)
		if number == 5 {
			printStatistics(cor, wro)
			fmt.Print("Choose operation:")
			continue
		} else if number == 0 {
			printStatistics(cor, wro)
			fmt.Println("Goodbye!")
			break
		}
		num := readInt(number)
		cal := calculate(num, a, b)
		che := checkAnswer(cal, userAnswer)
		if che == true {
			fmt.Println("Correct!")
		} else {
			fmt.Printf("Wrong! Correct answer: %v\n", cal)
		}
		fmt.Print("\nChoose operation:")
	}
}

func printMenu() {
	if number == 0 {
		return
	}
	fmt.Println("Menu")
	fmt.Print("1 - Addition\n2 - Subtraction\n3 - Multiplication\n4 - Integer division\n5 - Show statistics\n0 - Exit\n")
	fmt.Print("Choose operation:")
}

func readInt(message int) int {
	fmt.Printf("Enter first number: ")
	fmt.Scan(&a)
	fmt.Printf("Enter second number: ")
	fmt.Scan(&b)
	fmt.Printf("Enter your answer: ")
	fmt.Scan(&userAnswer)

	return message
}

func calculate(operation int, a int, b int) int {
	if operation == 1 {
		return a + b
	} else if operation == 2 {
		return a - b
	} else if operation == 3 {
		return a * b
	} else if operation == 4 {
		return a / b
	}
	return 0
}

func checkAnswer(correctAnswer int, userAns int) bool {
	if correctAnswer == userAns {
		cor++
		return true
	}
	wro++
	return false
}

func printStatistics(correct int, wrong int) {
	fmt.Println("Correct: ", correct)
	fmt.Println("Wrong: ", wrong)
}

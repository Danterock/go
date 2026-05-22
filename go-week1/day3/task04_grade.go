package main

import "fmt"

func main() {
	var score int
	fmt.Print("Enter your score ")
	fmt.Scan(&score)
	if score < 0 || score > 100 {
		fmt.Println("Invalid score")
	} else if score >= 0 && score <= 59 {
		fmt.Println("Grade: 2")
	} else if score >= 60 && score <= 74 {
		fmt.Println("Grade: 3")
	} else if score >= 75 && score <= 89 {
		fmt.Println("Grade: 4")
	} else if score >= 90 && score <= 100 {
		fmt.Println("Grade: 5")
	}
}

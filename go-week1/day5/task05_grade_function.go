package main

import "fmt"

func main() {
	var score int
	fmt.Print("Enter a score: ")
	fmt.Scan(&score)
	fmt.Println(getGrade(score))
}

func getGrade(score int) string {
	if score < 0 || score > 100 {
		return "Invalid score"
	} else if score >= 0 && score <= 59 {
		return "Grade: 2"
	} else if score >= 60 && score <= 74 {
		return "Grade: 3"
	} else if score >= 75 && score <= 89 {
		return "Grade: 4"
	} else if score >= 90 && score <= 100 {
		return "Grade: 5"
	}
	return "situation"
}

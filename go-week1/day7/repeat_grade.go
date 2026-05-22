package main

import "fmt"

func main() {
	var score int
	fmt.Printf("Enter score number: ")
	fmt.Scan(&score)
	if score < 0 || score > 100 {
		fmt.Println("Input out of range")
	} else if score > 0 && score <= 59 {
		fmt.Println(" Grade: 2")
	} else if score >= 60 && score <= 74 {
		fmt.Println(" Grade: 3")
	} else if score >= 75 && score <= 89 {
		fmt.Println(" Grade: 4")
	} else {
		fmt.Println(" Grade: 5")
	}

}

/*
90-100 -> Grade: 5
75-89 -> Grade: 4
60-74 -> Grade: 3
0-59 -> Grade: 2
less than 0 or more than 100 -> Invalid score
*/

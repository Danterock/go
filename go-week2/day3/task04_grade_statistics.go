package main

import "fmt"

func main() {
	var n int

	fmt.Print("input quantity: ")
	fmt.Scanln(&n)

	if n < 1 {
		fmt.Println("Invalid n")
		return
	}

	grades := make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Println("Input grade: ", i+1)
		fmt.Scan(&grades[i])

		if grades[i] < 2 || grades[i] > 5 {
			fmt.Println("Invalid grade")
			return
		}
	}
	//5 4 5 3 2 4 5 4
	stats := gradeStatistics(grades)
	for grade := 2; grade <= 5; grade++ {
		fmt.Printf("Grade %d: %d\n", grade, stats[grade])
	}
}
func gradeStatistics(grades []int) map[int]int {
	result := make(map[int]int)

	for i := 0; i < len(grades); i++ {
		result[grades[i]]++
	}
	return result
}

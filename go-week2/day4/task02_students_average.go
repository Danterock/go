package main

import (
	"fmt"
)

func main() {
	var quantity int
	fmt.Print("Input quantity of students: ")
	fmt.Scan(&quantity)
	students := make([]Student, quantity)
	if len(students) == 0 {
		fmt.Println("No average score")
		return
	}
	for i := 0; i < quantity; i++ {
		fmt.Scan(
			&students[i].Name,
			&students[i].Age,
			&students[i].Score,
		)
	}
	avg := averageScore(students)
	best := bestStudent(students)
	fmt.Printf("Average score: %.2f\n", avg)
	fmt.Println("Best student:", best.Name)
}
func averageScore(students []Student) float64 {
	sum := 0

	for i := 0; i < len(students); i++ {
		sum += students[i].Score
	}

	return float64(sum) / float64(len(students))
}
func bestStudent(students []Student) Student {
	if len(students) == 0 {
		return Student{}
	}
	best := students[0]
	for i := 1; i < len(students); i++ {
		if students[i].Score > best.Score {
			best = students[i]
		}
	}

	return best
}

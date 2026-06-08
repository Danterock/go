package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Score int
}

func main() {
	var student Student

	fmt.Scan(&student.Name, &student.Age, &student.Score)

	if student.Age <= 0 || student.Score < 0 || student.Score > 100 {
		fmt.Println("Invalid student")
		return
	}

	printStudent(student)
}

func printStudent(student Student) {
	fmt.Println("Name:", student.Name)
	fmt.Println("Age:", student.Age)
	fmt.Println("Score:", student.Score)
	fmt.Println("Passed:", student.Score >= 60)
}

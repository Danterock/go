package main

import "fmt"

type Student struct {
	Name   string
	Scores []int
}

func (s Student) Average() float64 {
	result := 0
	if len(s.Scores) < 0 {
		fmt.Println("Invalid scores")
	}
	for i := 0; i < len(s.Scores); i++ {
		result += s.Scores[i]
	}
	return float64(result) / float64(len(s.Scores))
}
func (s Student) Passed() bool {
	return s.Average() >= 60
}

func main() {
	var student Student
	var n int
	fmt.Print("Введите имя, колчиство оценок, оценки:")
	fmt.Scan(&student.Name)
	fmt.Scan(&n)
	if n < 1 {
		fmt.Println("Invalid scores")
		return
	}
	student.Scores = make([]int, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&student.Scores[i])
	}
	fmt.Printf("Average: %.2f\n", student.Average())
	fmt.Printf("Passed: %t\n", student.Passed())

}

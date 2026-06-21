package stats

import (
	"fmt"
	"week3-tasker/task"
)

func CountDone(tasks []task.Task) int {
	count := 0

	if len(tasks) <= 0 {
		return 0
	}

	for _, task := range tasks {
		if task.Done {
			count++
		}
	}
	return count
}

func CountActive(tasks []task.Task) int {
	count := 0

	if len(tasks) <= 0 {
		return 0
	}
	for _, task := range tasks {
		if !task.Done {
			count++
		}
	}
	return count
}

func CountByPriority(tasks []task.Task, priority string) int {
	count := 0

	if len(tasks) <= 0 {
		return 0
	}
	for _, task := range tasks {
		if task.Priority == priority {
			count++
		}
	}
	return count
}
func PrintStat(tasks []task.Task) {
	fmt.Println("Выполнено:", CountDone(tasks))
	fmt.Println("Активно:", CountActive(tasks))
	fmt.Println("High:", CountByPriority(tasks, "high"))
	fmt.Println("Medium:", CountByPriority(tasks, "medium"))
	fmt.Println("Low:", CountByPriority(tasks, "low"))
}

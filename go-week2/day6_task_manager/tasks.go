package main

import (
	"fmt"
	"strings"
)

func addTask(tasks []Task, nextID int, title string, priority string) ([]Task, int, bool) {
	title = strings.TrimSpace(title)

	if title == "" {
		return tasks, nextID, false
	}

	priority, ok := normalizePriority(priority)

	if !ok {
		return tasks, nextID, false
	}
	task := Task{
		ID:       nextID,
		Title:    title,
		Priority: priority,
		Done:     false,
	}

	tasks = append(tasks, task)

	nextID++

	return tasks, nextID, true
}

func findTaskIndexByID(tasks []Task, id int) int {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			return i
		}
	}
	return -1
}

func filterActiveTasks(tasks []Task) []Task {
	result := []Task{}
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Done == false {
			result = append(result, tasks[i])
		}
	}
	return result
}

func markDone(tasks []Task, id int) bool {
	index := findTaskIndexByID(tasks, id)
	if index == -1 {
		return false
	}
	tasks[index].Done = true
	return true
}

func deleteTask(tasks []Task, id int) ([]Task, bool) {
	index := findTaskIndexByID(tasks, id)
	if index == -1 {
		return tasks, false
	}
	tasks = append(tasks[:index], tasks[index+1:]...)
	return tasks, true
}

func searchTasks(tasks []Task, query string) []Task {
	result := []Task{}
	query = strings.ToLower(query)
	for i := 0; i < len(tasks); i++ {
		title := strings.ToLower(tasks[i].Title)
		if strings.Contains(title, query) {
			result = append(result, tasks[i])
		}
	}
	return result
}

func getStatistics(tasks []Task) map[string]int {
	stats := map[string]int{
		"total":  0,
		"active": 0,
		"done":   0,
		"high":   0,
		"medium": 0,
		"low":    0,
	}
	stats["total"] = len(tasks)
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Done == false {
			stats["active"]++
		} else {
			stats["done"]++
		}
		switch tasks[i].Priority {
		case "high":
			stats["high"]++

		case "medium":
			stats["medium"]++

		case "low":
			stats["low"]++
		default:
			break
		}
	}
	return stats
}

func printTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("no tasks")
		return
	}
	for i := 0; i < len(tasks); i++ {
		status := "active"
		if tasks[i].Done {
			status = "done"
		}
		fmt.Printf(
			"[%d] %s | %s | %s\n",
			tasks[i].ID,
			tasks[i].Title,
			tasks[i].Priority,
			status,
		)
	}
}

func printMenu() {
	fmt.Print("1 - Add task\n" +
		"2 - Show all tasks\n" +
		"3 - Show active tasks\n" +
		"4 - Mark task as done\n" +
		"5 - Delete task\n" +
		"6 - Search by title\n" +
		"7 - Show statistics\n" +
		"0 - Exit")

}

func printStats(stats map[string]int) {
	fmt.Println("Total:", stats["total"])
	fmt.Println("Active:", stats["active"])
	fmt.Println("Done:", stats["done"])
	fmt.Println("High:", stats["high"])
	fmt.Println("Medium:", stats["medium"])
	fmt.Println("Low:", stats["low"])
}

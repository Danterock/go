package main

import (
	"fmt"
)

func main() {
	tasks := []Task{}
	nextID := 1
	ok := false
	for {
		printMenu()
		inputNumber := readInt("\nChoose an option: ")
		switch inputNumber {
		case 1:
			title := readLine("Title: ")
			priority := readLine("Priority: ")
			tasks, nextID, ok = addTask(tasks, nextID, title, priority)
			if !ok {
				if title == "" {
					fmt.Println("Empty title is not allowed")
				} else {
					fmt.Println("Invalid priority")
				}
			}
		case 2:
			printTasks(tasks)
		case 3:
			active := filterActiveTasks(tasks)
			printTasks(active)
		case 4:
			id := readInt("ID: ")
			if !markDone(tasks, id) {
				fmt.Println("Task not found")
			}
		case 5:
			id := readInt("ID: ")
			tasks, ok = deleteTask(tasks, id)
			if !ok {
				fmt.Println("Task not found")
			}
		case 6:
			query := readLine("Input title: ")
			found := searchTasks(tasks, query)
			printTasks(found)
		case 7:
			stats := getStatistics(tasks)
			printStats(stats)
		case 0:
			stats := getStatistics(tasks)
			printStats(stats)
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}

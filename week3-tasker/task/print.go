package task

import (
	"fmt"
)

func PrintTasks(tasks []Task) {
	for _, t := range tasks {
		fmt.Printf("%d | %s | %s | %t\n",
			t.ID,
			t.Title,
			t.Priority,
			t.Done,
		)
	}
}
func PrintMenu() {
	fmt.Print("1. Show all tasks\n" +
		"2. Add task\n" +
		"3. Show active tasks\n" +
		"4. Mark task as done\n" +
		"5. Delete task by ID\n" +
		"6. Search task by title\n" +
		"7. Show statistics\n" +
		"8. Save tasks\n" +
		"0. Exit\n")
}
func PrintActiveTask(tasks []Task) {
	result := ActiveTasks(tasks)
	for i := 0; i < len(result); i++ {
		fmt.Printf("%d | %s | %s | %t\n", result[i].ID, result[i].Title, result[i].Priority, result[i].Done)
	}
}

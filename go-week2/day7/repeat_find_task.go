package main

import "fmt"

type Task struct {
	ID    int
	Title string
	Done  bool
}

func main() {
	tasks := []Task{
		{1, "Learn slices", false},
		{2, "Practice strings", true},
		{3, "Build task manager", false},
	}

	var id int
	fmt.Print("Input task ID: ")
	fmt.Scan(&id)
	find, ok := findTaskByID(tasks, id)

	if !ok {
		fmt.Println("Not found.")
	}
	fmt.Println("Task:", find.Title)
	fmt.Println("Done:", find.Done)
}

func findTaskByID(tasks []Task, id int) (Task, bool) {
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			return tasks[i], true
		}
	}
	return Task{}, false
}

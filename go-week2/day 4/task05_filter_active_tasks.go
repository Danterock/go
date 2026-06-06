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
	filter := filterActiveTasks(tasks)

	if len(filter) == 0 {
		fmt.Println("No active tasks")
	}

	for i := 0; i < len(filter); i++ {
		fmt.Printf("%d %s %v\n", filter[i].ID, filter[i].Title, filter[i].Done)
	}

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

package main

import (
	"fmt"
	"path/filepath"
	"week3-tasker/stats"
	"week3-tasker/storage"
	"week3-tasker/task"
)

func main() {
	var tasks []task.Task
	var err error
	tasksPath := filepath.Join("data", "tasks.json")
	reportPath := filepath.Join("data", "report.txt")
	logPath := filepath.Join("data", "app.log")
	tasks, err = storage.LoadTasksJSON(tasksPath)
	if err != nil {
		fmt.Println("Open JSON error: ", err)
		return
	}
	defer storage.SaveTasksJSON(tasksPath, tasks)
	for {
		task.PrintMenu()
		switch task.ReadInt("Choose operation: ") {
		case 0:
			return
		case 1:
			task.PrintTasks(tasks)
		case 2:
			tasks, err = task.AddTask(tasks, task.ReadLine("input title"), task.ReadLine("input priority"))
			if err != nil {
				fmt.Println("Add Task error: ", err)
			}
			newTask := tasks[len(tasks)-1]
			err = storage.AppendLog(logPath, fmt.Sprintf("action=add id=%d title=%s priority=%s", newTask.ID, newTask.Title, newTask.Priority))
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			task.PrintActiveTask(tasks)
		case 4:
			tasks, err = task.MarkDone(tasks, task.ReadInt("Input id to mark done: "))
			if err != nil {
				fmt.Println("Mark Done error: ", err)
			}
		case 5:
			tasks, err = task.DeleteByID(tasks, task.ReadInt("Input id to delete: "))
			if err != nil {
				fmt.Println("Delete Task error: ", err)
			}
		case 6:
			found, err := task.SearchByTitle(tasks, task.ReadLine("input title: "))
			task.PrintTasks(found)
			if err != nil {
				fmt.Println("Search Task error: ", err)
			}
		case 7:
			stats.PrintStat(tasks)
		case 8:
			err = storage.SaveTasksJSON(tasksPath, tasks)
			if err != nil {
				fmt.Println("Save JSON error: ", err)
			}
		default:
			fmt.Println("Invalid operation")
		}

		err = storage.SaveReport(reportPath, tasks)
		if err != nil {
			fmt.Println(err)
		}
	}
}

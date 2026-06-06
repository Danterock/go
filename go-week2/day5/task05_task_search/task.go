package task05_task_search

type Task struct {
	ID    int
	Title string
	Done  bool
}

func FilterActiveTasks(tasks []Task) []Task {
	newTasks := []Task{}
	if len(tasks) == 0 {
		return newTasks
	}
	for _, task := range tasks {
		if task.Done == false {
			newTasks = append(newTasks, task)
		}
	}
	return newTasks
}

func FindTaskByID(tasks []Task, id int) (Task, bool) {
	newTask := Task{}
	if len(tasks) == 0 {
		return newTask, false
	}
	for _, task := range tasks {
		if task.ID == id {
			newTask = task
		}
	}
	return newTask, true
}

package task05_task_search

import "testing"

func TestFindTaskByID(t *testing.T) {
	tasks := []Task{
		{1, "Learn Go", false},
		{2, "Practice", true},
	}

	task, ok := FindTaskByID(tasks, 1)

	if task.ID != 1 {
		t.Error("Задача должна была быть найдена")
	}
	if !ok {
		t.Error("должно быть тру")
	}
}
func TestNoID(t *testing.T) {
	tasks := []Task{}

	_, ok := FindTaskByID(tasks, 0)
	if ok {
		t.Error("должно быть фолс")
	}
	result := FilterActiveTasks(tasks)

	if len(result) != 0 {
		t.Error("должно быть 0")
	}
}

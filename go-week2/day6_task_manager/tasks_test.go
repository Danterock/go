package main

import "testing"

func TestFindTaskIndexByID(t *testing.T) {
	tasks := []Task{
		{1, "A", "high", false},
		{2, "B", "low", false},
	}
	index := findTaskIndexByID(tasks, 2)

	if index != 1 {
		t.Error("ожидали 1")
	}
}
func TestFilterActiveTasks(t *testing.T) {
	tasks := []Task{
		{1, "A", "high", false},
		{2, "B", "low", false},
		{3, "C", "Medium", true},
	}
	result := filterActiveTasks(tasks)

	if len(result) != 2 {
		t.Error("должно быть 2")
	}
}
func TestMarkDone(t *testing.T) {
	tasks := []Task{
		{1, "A", "high", false},
		{2, "B", "low", true},
		{3, "C", "Medium", false},
	}
	result := markDone(tasks, 1)
	result1 := markDone(tasks, -1)

	if result != true {
		t.Error("должно быть true")
	}
	if result1 != false {
		t.Error("должно быть false")
	}
}
func TestDeleteTask(t *testing.T) {
	tasks := []Task{
		{1, "A", "high", false},
		{2, "B", "low", true},
		{3, "C", "Medium", false},
	}
	_, ok := deleteTask(tasks, 1)
	if ok != true {
		t.Error("ожидали true")
	}
}
func TestSearchTasks(t *testing.T) {
	tasks := []Task{
		{1, "A", "high", false},
		{2, "B", "low", true},
		{3, "C", "Medium", false},
	}
	result := searchTasks(tasks, "A")
	if len(result) != 1 {
		t.Error("а должно быть 1")
	}
}

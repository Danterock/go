package task

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func AddTask(tasks []Task, title, priority string) ([]Task, error) {
	err := ValidateTitle(title)
	if err != nil {
		return tasks, err
	}
	priority, err = ValidatePriority(priority)
	if err != nil {
		return tasks, err
	}
	newTasks := Task{
		Title:    title,
		Priority: priority,
		Done:     false,
		ID:       NextID(tasks),
	}
	tasks = append(tasks, newTasks)
	return tasks, nil
}
func MarkDone(tasks []Task, id int) ([]Task, error) {
	_, err := FindByID(tasks, id)
	if err != nil {
		return tasks, err
	}

	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Done = true
			return tasks, nil
		}
	}

	return tasks, errors.New("task not found")
}
func ActiveTasks(tasks []Task) []Task {
	result := []Task{}
	for i := 0; i < len(tasks); i++ {
		if tasks[i].Done == false {
			result = append(result, tasks[i])
		}
	}
	return result
}
func DeleteByID(tasks []Task, id int) ([]Task, error) {
	_, err := FindByID(tasks, id)
	if err != nil {
		return tasks, err
	}
	for i := 0; i < len(tasks); i++ {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
		}
	}
	return tasks, nil
}
func ReadLine(message string) string {
	fmt.Println(message)
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	return strings.TrimSpace(text)
}

func ReadInt(message string) int {
	text := ReadLine(message)

	number, err := strconv.Atoi(text)
	if err != nil {
		return -1
	}

	return number
}
func SearchByTitle(tasks []Task, query string) ([]Task, error) {
	result := []Task{}
	query = strings.ToLower(query)
	for i := 0; i < len(tasks); i++ {
		title := strings.ToLower(tasks[i].Title)
		if strings.Contains(title, query) {
			result = append(result, tasks[i])
		}
	}
	if len(result) == 0 {
		return result, errors.New("task not found")
	}
	if len(result) > len(tasks) {
		return result, errors.New("too many tasks")
	}
	return result, nil
}
func NextID(tasks []Task) int {
	maxID := 0
	for i := range tasks {
		if tasks[i].ID > maxID {
			maxID = tasks[i].ID
		}
	}
	return maxID + 1
}

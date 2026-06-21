package task

import (
	"errors"
	"strings"
)

func ValidateTitle(input string) error {
	input = strings.TrimSpace(input)
	if input == "" {
		return errors.New("Empty title")
	}
	return nil
}

func ValidatePriority(priority string) (string, error) {
	priority = strings.TrimSpace(priority)
	if priority == "" {
		return "medium", nil
	}
	if priority != "low" && priority != "medium" && priority != "high" {
		return "", errors.New("Invalid priority")
	}
	return priority, nil
}

func FindByID(tasks []Task, id int) (Task, error) {
	for _, task := range tasks {
		if task.ID == id {
			return task, nil
		}
	}
	return Task{}, errors.New("Task not found")
}

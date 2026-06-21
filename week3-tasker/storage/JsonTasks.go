package storage

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"week3-tasker/stats"
	"week3-tasker/task"
)

func SaveTasksJSON(path string, tasks []task.Task) error {
	data, err := json.MarshalIndent(tasks, "", "    ")
	if err != nil {
		return err
	}
	err1 := os.WriteFile(path, data, 0644)
	if err1 != nil {
		return errors.New("Can't save tasks.")
	}
	return nil
}
func LoadTasksJSON(path string) ([]task.Task, error) {
	var tasks []task.Task
	data, err := os.ReadFile(path)
	if os.IsNotExist(err) {
		return tasks, nil
	}
	if err != nil {
		return nil, err
	}
	if len(data) == 0 {
		return tasks, nil
	}
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		return nil, err
	}
	for i := range tasks {
		if tasks[i].Priority == "" {
			tasks[i].Priority = "medium"
		}
	}
	return tasks, nil
}

func SaveReport(path string, tasks []task.Task) error {
	report := fmt.Sprintf(
		"Всего задач: %d\nВыполнено: %d\nАктивно: %d\nHigh: %d\nMedium: %d\nLow: %d\n",
		len(tasks),
		stats.CountDone(tasks),
		stats.CountActive(tasks),
		stats.CountByPriority(tasks, "high"),
		stats.CountByPriority(tasks, "medium"),
		stats.CountByPriority(tasks, "low"),
	)

	return os.WriteFile(path, []byte(report), 0644)
}
func AppendLog(path, action string) error {
	file, err := os.OpenFile(path, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(action + "\n")
	return err
}

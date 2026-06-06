package main

import (
	"bufio"
	"fmt"
	"os"
)

type Task struct {
	ID    int
	Title string
	Done  bool
}

func (t Task) IsDone() bool {
	return t.Done
}
func (t *Task) MarkDone() {
	t.Done = true
}

func main() {
	var task Task

	fmt.Scanln(&task.ID)

	reader := bufio.NewReader(os.Stdin)
	task.Title, _ = reader.ReadString('\n')

	if task.IsDone() {
		fmt.Println("Before: done")
	} else {
		fmt.Println("Before: active")
	}

	task.MarkDone()

	if task.IsDone() {
		fmt.Println("After: done")
	} else {
		fmt.Println("After: active")

	}
}
